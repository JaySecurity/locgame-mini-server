package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/jobs"
)

type Jobs struct {
	db     *mongo.Client
	config *config.Config

	jobs          *mongo.Collection
	recurringJobs *mongo.Collection
}

// NewJobs creates a new instance of the distributed locks store.
func NewJobs(config *config.Config, db *mongo.Client) *Jobs {
	s := new(Jobs)
	s.db = db
	s.config = config

	s.jobs = db.Database(s.config.Database.Database).Collection("jobs")
	s.recurringJobs = db.Database(s.config.Database.Database).Collection("recurring_jobs")

	return s
}

func (s *Jobs) SaveRecurringJob(ctx context.Context, job *jobs.RecurringJobData) error {
	opts := options.Update().SetUpsert(true)
	_, err := s.recurringJobs.UpdateByID(ctx, job.ID, bson.M{"$set": job}, opts)
	return err
}

func (s *Jobs) SaveJob(ctx context.Context, job *jobs.JobData) error {
	opts := options.Update().SetUpsert(true)
	_, err := s.jobs.UpdateByID(ctx, job.ID, bson.M{"$set": job}, opts)
	return err
}

func (s *Jobs) DeleteJob(ctx context.Context, id *base.ObjectID) error {
	_, err := s.recurringJobs.DeleteOne(ctx, bson.M{"_id": id})
	if err == nil {
		_, err = s.jobs.DeleteMany(ctx, bson.M{"parent_job_id": id})
	}
	return err
}

func (s *Jobs) GetRecurringJobs(ctx context.Context) (map[string]*jobs.RecurringJobData, error) {
	cursor, err := s.recurringJobs.Find(ctx, bson.M{})
	if err == nil {
		recurringJobs := make(map[string]*jobs.RecurringJobData)
		for cursor.Next(ctx) {
			var job jobs.RecurringJobData
			_ = cursor.Decode(&job)
			recurringJobs[job.Name] = &job
		}
		_ = cursor.Close(ctx)
		return recurringJobs, err
	}
	return nil, err
}
