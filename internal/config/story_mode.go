package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	gameDto "locgame-mini-server/pkg/dto/game"
)

func init() {
	Register(NewStoryModeConfig)
}

// StoryModeConfig stores story mode configuration.
type StoryModeConfig struct {
	BaseConfig

	Missions map[int32]*gameDto.StoryModeMissionData `yaml:"Levels"`
}

// NewStoryModeConfig creates an instance of the story mode configuration.
func NewStoryModeConfig() map[int32]*gameDto.StoryModeMissionData {
	c := new(StoryModeConfig)
	c.self = c
	c.Load("story_mode.yaml")
	return c.Missions
}

func (c *StoryModeConfig) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	if err == nil {
		err = yaml.Unmarshal(bytes, &c)
	}

	return err
}
