package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"locgame-mini-server/pkg/dto/resources"
)

func init() {
	Register(NewDailyRewards)
}

// DailyRewards stores daily rewards configuration.
type DailyRewards struct {
	BaseConfig

	Rewards []*resources.ResourceAdjustment
}

// NewDailyRewards creates an instance of the daily rewards configuration.
func NewDailyRewards() []*resources.ResourceAdjustment {
	c := new(DailyRewards)
	c.self = c
	c.Load("daily_rewards.yaml")
	return c.Rewards
}

func (c *DailyRewards) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	if err == nil {
		err = yaml.Unmarshal(bytes, &c.Rewards)
	}

	return err
}
