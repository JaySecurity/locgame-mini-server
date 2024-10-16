package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func init() {
	Register(NewWithdrawConfig)
}

// WithdrawConfig stores withdraw configuration.
type WithdrawConfig struct {
	BaseConfig

	Min    int32   `yaml:"Min"`
	Max    int32   `yaml:"Max"`
	MinFee float64 `yaml:"MinFee"`
	MaxFee float64 `yaml:"MaxFee"`
}

// NewWithdrawConfig creates an instance of the withdraw configuration.
func NewWithdrawConfig() *WithdrawConfig {
	c := new(WithdrawConfig)
	c.self = c
	c.Load("withdraw.yaml")
	return c
}

func (c *WithdrawConfig) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	if err == nil {
		err = yaml.Unmarshal(bytes, &c)
	}

	return err
}
