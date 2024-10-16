package jobs

import (
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/store"
)

type Job interface {
	Run() error
	Init(config *config.Config, store *store.Store, logger *JobLogger)
	GetConfig() *config.Config
	GetStore() *store.Store
	GetLogger() *JobLogger

	GetPrefix() string
	GetName() string
	GetRetries() uint32

	ApplyOptions(options ...JobOption)
}

type BaseJob struct {
	config *config.Config
	store  *store.Store
	logger *JobLogger

	prefix  string
	retries uint32
	name    string
}

func (j *BaseJob) Init(config *config.Config, store *store.Store, logger *JobLogger) {
	j.config = config
	j.store = store
	j.logger = logger
}

func (j *BaseJob) Run() error {
	j.logger.Warning("Incorrect job")
	return nil
}

func (j *BaseJob) GetConfig() *config.Config {
	return j.config
}

func (j *BaseJob) GetStore() *store.Store {
	return j.store
}

func (j *BaseJob) GetLogger() *JobLogger {
	return j.logger
}

func (j *BaseJob) GetName() string {
	return j.name
}

func (j *BaseJob) GetPrefix() string {
	return j.prefix
}

func (j *BaseJob) GetRetries() uint32 {
	return j.retries
}

func (j *BaseJob) ApplyOptions(options ...JobOption) {
	for _, option := range options {
		option(j)
	}
}

type JobOption func(job *BaseJob)

func WithPrefix(prefix string) JobOption {
	return func(job *BaseJob) {
		job.prefix = prefix
	}
}

func WithRetry(retries uint32) JobOption {
	return func(job *BaseJob) {
		job.retries = retries
	}
}

func WithName(name string) JobOption {
	return func(job *BaseJob) {
		job.name = name
	}
}
