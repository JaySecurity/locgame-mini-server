package jobs

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/jobs"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/pubsub"
)

var (
	cronParser = cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	cronJobs   = cron.New(cron.WithParser(cronParser), cron.WithLocation(time.UTC))

	cfg           *config.Config
	db            *store.Store
	recurringJobs map[string]*recurringJob
	initializers  []func(config *config.Config, store *store.Store)
	workers       []Worker
)

type recurringJob struct {
	job     Job
	data    *jobs.RecurringJobData
	jobFunc cron.FuncJob
}

func init() {
	recurringJobs = make(map[string]*recurringJob)
}

func Run() {
	config.SetShowLog(false)
	cfg = config.Init("configs/")
	db = store.NewStore(cfg)

	cfg.OnReloadComplete = append(cfg.OnReloadComplete, func() {
		for _, initializer := range initializers {
			initializer(cfg, db)
		}

		addAllRegisteredJobs()
	})

	for _, worker := range workers {
		worker.Init(cfg, db, NewJobLogger())
	}

	pubsub.RegisterHandler(&TriggerNowHandler{})

	pubsub.Connect(nil, cfg.NatsAddress)

	rand.Seed(time.Now().UTC().Unix())

	for _, initializer := range initializers {
		initializer(cfg, db)
	}

	addAllRegisteredJobs()

	log.Debug("All cron jobs registered")
	log.Debug("Service is running...")

	for _, worker := range workers {
		go worker.Run()
	}

	cronJobs.Run()
}

func Stop() {
	cronJobs.Stop()

	wg := sync.WaitGroup{}
	for _, worker := range workers {
		wg.Add(1)
		w := worker
		go func() {
			w.Stop()
			wg.Done()
		}()
	}
	wg.Wait()
}

func addAllRegisteredJobs() {
	clearJobs()

	ctx := context.Background()
	recurringJobsData, err := db.Jobs.GetRecurringJobs(ctx)
	if err != nil {
		log.Fatal("Unable to get recurring jobs:", err)
	}
	for _, recurringJob := range recurringJobs {
		if recurringJob.data.Disabled {
			continue
		}

		if jobData, ok := recurringJobsData[recurringJob.data.Name]; ok {
			if recurringJob.data.Schedule != jobData.Schedule || recurringJob.data.Retries != jobData.Retries {
				jobData.Schedule = recurringJob.data.Schedule
				jobData.Retries = recurringJob.data.Retries
				_ = db.Jobs.SaveRecurringJob(ctx, jobData)
			}
			recurringJob.data = jobData
		} else {
			recurringJob.data.ID = &base.ObjectID{Value: primitive.NewObjectID().Hex()}
			_ = db.Jobs.SaveRecurringJob(ctx, recurringJob.data)
		}
		err = addJob(recurringJob)
		if err != nil {
			log.Fatal(err)
		}
	}

	for id, data := range recurringJobsData {
		if _, exists := recurringJobs[id]; !exists {
			err = db.Jobs.DeleteJob(ctx, data.ID)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func addJob(recurringJob *recurringJob) error {
	_, err := cronJobs.AddJob(recurringJob.data.Schedule, recurringJob.jobFunc)
	return err
}

func initJob(parentJobID *base.ObjectID, name string, job Job) (Job, *jobs.JobData) {
	job.Init(cfg, db, NewJobLogger())

	data := &jobs.JobData{
		ID:          &base.ObjectID{Value: primitive.NewObjectID().Hex()},
		ParentJobID: parentJobID,
		Name:        name,
		Status:      jobs.JobStatus_Running,
		StartedAt:   &base.Timestamp{Seconds: time.Now().UTC().Unix()},
	}

	return job, data
}

func OnInit(fn func(config *config.Config, store *store.Store)) {
	initializers = append(initializers, fn)
}

func RegisterWorker(worker Worker) {
	workers = append(workers, worker)
}

func Register(schedule string, job Job, options ...JobOption) {
	job.ApplyOptions(options...)
	name := job.GetName()
	if name == "" {
		name = job.GetPrefix() + ToSnakeCase(strings.Split(fmt.Sprintf("%v", reflect.TypeOf(job)), ".")[1])
	}
	recurringJob := &recurringJob{
		data: &jobs.RecurringJobData{
			Name:     name,
			Schedule: schedule,
			Retries:  job.GetRetries(),
			Status:   jobs.JobStatus_NotSet,
		},
		job: job,
	}
	recurringJob.jobFunc = getJobFunc(context.Background(), recurringJob)
	recurringJobs[name] = recurringJob
}

func getJobFunc(ctx context.Context, recurringJob *recurringJob) cron.FuncJob {
	return func() {
		// Do not set TTL duration less than timeout! Otherwise it will be ContextDeadlineExceeded.
		mutex := db.DistributedLocks.NewLock(recurringJob.data.Name, 8*time.Second)
		if err := mutex.LockContext(ctx); err != nil {
			return
		}

		// Unlock will happen automatically after TTL expires

		recurringJob.data.LastExecution = &base.Timestamp{Seconds: time.Now().UTC().Unix()}
		recurringJob.data.Status = jobs.JobStatus_Running
		err := db.Jobs.SaveRecurringJob(ctx, recurringJob.data)

		if err != nil {
			log.Error(err)
		}

		job, data := initJob(recurringJob.data.ID, recurringJob.data.Name, recurringJob.job)
		err = db.Jobs.SaveJob(ctx, data)
		if err != nil {
			log.Error(err)
		}

		attempts := -1
		var jobErr error

		defer func() {
			data.Output = job.GetLogger().output.String()
			job.GetLogger().output.Reset()

			data.Attempt = uint32(attempts + 1)

			if jobErr == nil {
				data.Status = jobs.JobStatus_Success
			} else {
				data.Status = jobs.JobStatus_Failed
			}
			data.FinishedAt = &base.Timestamp{Seconds: time.Now().UTC().Unix()}

			err = db.Jobs.SaveJob(ctx, data)
			if err != nil {
				log.Error(err)
			}

			recurringJob.data.Status = data.Status
			err = db.Jobs.SaveRecurringJob(ctx, recurringJob.data)

			if err != nil {
				log.Error(err)
			}
		}()

		for attempts < int(recurringJob.data.Retries) {
			func() {
				defer func() {
					if r := recover(); r != nil {
						const size = 64 << 10
						buf := make([]byte, size)
						buf = buf[:runtime.Stack(buf, false)]
						err, ok := r.(error)
						if !ok {
							err = fmt.Errorf("%v", r)
						}
						job.GetLogger().Errorf("Panic: %v\n%s", err, buf)
					}
				}()
				jobErr = job.Run()
				if jobErr != nil {
					job.GetLogger().Error(jobErr)
				}
			}()
			if jobErr == nil || recurringJob.data.Retries == 0 {
				break
			} else {
				attempts++
				job.GetLogger().Warningf("The job returned an error. Repeat. Attempt %d of %d", attempts, recurringJob.data.Retries)
			}
		}
	}
}

func clearJobs() {
	for _, entry := range cronJobs.Entries() {
		cronJobs.Remove(entry.ID)
	}
}

func triggerNow(jobName string) {
	if job, ok := recurringJobs[jobName]; ok {
		job.jobFunc()
	}
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z\\d])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
