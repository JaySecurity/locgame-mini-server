package metrics

import (
	"locgame-mini-server/internal/config"
	"time"
)

type Metric interface {
	Log(message string)
	Error(message error)
	LogGameTimeAndMode(user interface{}, opponent interface{}, gameTime string,
		gameMode string, gameResult string, gameReward interface{}, progress ...int32)

	LCEarned(user interface{}, lc int32, reason ...string)
	LCSpent(user interface{}, lc int32, reason ...string)
	//LCConverted()
	LastLogin(user interface{}, time time.Time)
	LogGameMove(valid string, gameId string, playerId string, movement string, moveResult string, err ...error)
	//DropConnections()
	//- [ ]  match request made vs accepted vs no one available to play
}
type MultiMetric struct {
	instances []Metric
}

const FILE string = "file"
const OPENSEARCH string = "opensearch"

var defaultInstance *MultiMetric

func (ml *MultiMetric) LogGameTimeAndMode(user interface{}, opponent interface{}, gameTime string,
	gameMode string, gameResult string, gameReward interface{}, progress ...int32) {

	for _, logger := range ml.instances {
		logger.LogGameTimeAndMode(user, opponent, gameTime,
			gameMode, gameResult, gameReward, progress...)
	}
}

func (ml *MultiMetric) Log(message string) {
	done := make(chan struct{}, len(ml.instances))

	// Send the log message to each logger through a separate channel
	for _, logger := range ml.instances {
		logChan := make(chan string)
		go func(logger Metric, logChan chan string) {
			logger.Log(<-logChan)
			done <- struct{}{}
		}(logger, logChan)

		logChan <- message
	}

	// Wait for all log messages to be processed by each logger
	for i := 0; i < len(ml.instances); i++ {
		<-done
	}
}

func (ml *MultiMetric) Error(message error) {
	done := make(chan struct{}, len(ml.instances))

	// Send the log message to each logger through a separate channel
	for _, logger := range ml.instances {
		logChan := make(chan error)
		go func(logger Metric, logChan chan error) {
			logger.Error(<-logChan)
			done <- struct{}{}
		}(logger, logChan)

		logChan <- message
	}

	// Wait for all log messages to be processed by each logger
	for i := 0; i < len(ml.instances); i++ {
		<-done
	}
}

func (ml *MultiMetric) LCEarned(user interface{}, lc int32, reason ...string) {
	for _, logger := range ml.instances {
		logger.LCEarned(user, lc, reason...)
	}
}

func (ml *MultiMetric) LastLogin(user interface{}, lastLogin time.Time) {
	for _, logger := range ml.instances {
		logger.LastLogin(user, lastLogin)
	}
}

func (ml *MultiMetric) LCSpent(user interface{}, lc int32, reason ...string) {
	for _, logger := range ml.instances {
		logger.LCSpent(user, lc, reason...)
	}
}

func (ml *MultiMetric) LogGameMove(valid string, gameId string, playerId string, movement string, moveResult string, err ...error) {
	for _, logger := range ml.instances {
		logger.LogGameMove(valid, gameId, playerId, movement, moveResult, err...)
	}
}

func NewMetric(drivers ...string) Metric {
	metricInstances := newMultiMetricArray(nil, drivers...)
	return &MultiMetric{instances: metricInstances}
}

func newMultiMetricArray(cfg *config.Config, drivers ...string) []Metric {
	var metricInstances []Metric
	for _, driver := range drivers {
		switch driver {
		case FILE:
			metricInstances = append(metricInstances, NewFileLogger())
		case OPENSEARCH:
			metricInstances = append(metricInstances, NewZapLogger(cfg.Metrics.MetricsByID[OPENSEARCH]))
		default:
		}
	}

	return metricInstances
}

func GetDefault() Metric {
	if defaultInstance == nil {
		return SetDefault(nil, FILE)
	}
	return defaultInstance
}

func SetDefault(cfg *config.Config, drivers ...string) Metric {
	defaultInstance = &MultiMetric{
		instances: newMultiMetricArray(cfg, drivers...),
	}
	return defaultInstance
}
