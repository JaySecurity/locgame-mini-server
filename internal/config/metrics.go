package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func init() {
	Register(NewMetricsConfig)
}

type Metric struct {
	MetricID   string `yaml:"MetricID,omitempty"`
	APIAddress string `yaml:"APIAddress,omitempty"`
	Username   string `yaml:"Username,omitempty"`
	Password   string `yaml:"Password,omitempty"`
}

// MetricsConfig stores metrics configuration.
type MetricsConfig struct {
	BaseConfig

	allMetricsIDs []string
	MetricsByID   map[string]*Metric
}

// NewMetricsConfig creates an instance of the metrics' configuration.
func NewMetricsConfig() *MetricsConfig {
	c := new(MetricsConfig)
	c.self = c
	c.Load("metrics.yaml")
	return c
}

func (c *MetricsConfig) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	data := struct {
		Ms []*Metric `yaml:"Metrics"`
	}{}
	if err == nil {
		if err = yaml.Unmarshal(bytes, &data); err == nil {
			c.MetricsByID = make(map[string]*Metric)
			for _, metric := range data.Ms {
				c.MetricsByID[metric.MetricID] = metric
				c.allMetricsIDs = append(c.allMetricsIDs, metric.MetricID)
			}
		}
	}

	return err
}

func (c *MetricsConfig) GetAllMetrics() []string {
	return c.allMetricsIDs
}
