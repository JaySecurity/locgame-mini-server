package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"locgame-mini-server/pkg/dto/arena"
)

func init() {
	Register(NewLeaguesConfig)
}

// LeaguesConfig stores leagues configuration.
type LeaguesConfig struct {
	BaseConfig

	Data map[int32]*arena.LeagueData
}

// NewLeaguesConfig creates an instance of the leagues configuration.
func NewLeaguesConfig() map[int32]*arena.LeagueData {
	c := new(LeaguesConfig)
	c.self = c
	c.Load("leagues.yaml")
	return c.Data
}

func (c *LeaguesConfig) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	var data []*arena.LeagueData

	if err == nil {
		if err = yaml.Unmarshal(bytes, &data); err == nil {
			c.Data = make(map[int32]*arena.LeagueData, 0)
			for _, v := range data {
				c.Data[int32(v.Type)] = v
			}
		}
	}
	return err
}
