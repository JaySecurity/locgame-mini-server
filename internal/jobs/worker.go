package jobs

import (
	"reflect"
	"sync"

	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/pkg/pubsub"
)

type Worker interface {
	Run()
	Stop()
	Init(config *config.Config, store *store.Store, logger *JobLogger)
	GetConfig() *config.Config
	GetStore() *store.Store
	GetLogger() *JobLogger

	GetPrefix() string
	GetName() string
	GetRetries() uint32

	GetDataType() reflect.Type

	Handle(data pubsub.MessageData)
}

type BaseWorker struct {
	config *config.Config
	store  *store.Store
	logger *JobLogger

	prefix  string
	retries uint32
	name    string

	Done      chan struct{}
	WaitGroup sync.WaitGroup
}

func (j *BaseWorker) Init(config *config.Config, store *store.Store, logger *JobLogger) {
	j.config = config
	j.store = store
	j.logger = logger
}

func (j *BaseWorker) Stop() {
	j.Done <- struct{}{}
	j.WaitGroup.Wait()
}

func (j *BaseWorker) Run() error {
	j.logger.Warning("Incorrect worker")
	return nil
}

func (j *BaseWorker) GetConfig() *config.Config {
	return j.config
}

func (j *BaseWorker) GetStore() *store.Store {
	return j.store
}

func (j *BaseWorker) GetLogger() *JobLogger {
	return j.logger
}

func (j *BaseWorker) GetName() string {
	return j.name
}

func (j *BaseWorker) GetPrefix() string {
	return j.prefix
}

func (j *BaseWorker) GetRetries() uint32 {
	return j.retries
}
