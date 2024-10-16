package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"locgame-mini-server/pkg/dto/arena"
)

func init() {
	Register(NewTrophyRoadRewards)
}

// TrophyRoadRewards stores trophy road rewards configuration.
type TrophyRoadRewards struct {
	BaseConfig

	Data []*arena.TrophyRoadRewards
}

// NewTrophyRoadRewards creates an instance of the trophy road rewards configuration.
func NewTrophyRoadRewards() []*arena.TrophyRoadRewards {
	c := new(TrophyRoadRewards)
	c.self = c
	c.Load("trophy_road_rewards.yaml")
	return c.Data
}

func (c *TrophyRoadRewards) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	var data []*arena.TrophyRoadRewards

	if err == nil {
		if err = yaml.Unmarshal(bytes, &data); err == nil {
			c.Data = data
		}
	}

	return err
}
