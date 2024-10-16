package config

import (
	"io/ioutil"
	"locgame-mini-server/pkg/dto/game"

	"gopkg.in/yaml.v3"
)

func init() {
	Register(NewMatchTimerConfig)
}

// MatchTimerConfig stores MatchTimer configuration.
type MatchTimerConfig struct {
	BaseConfig

	Duration     int32 `yaml:"Duration"`
	StoryMode    bool  `yaml:"StoryMode"`
	ArenaMode    bool  `yaml:"ArenaMode"`
	FriendlyMode bool  `yaml:"FriendlyMode"`
	QuickMode    bool  `yaml:"QuickMode"`
}

// NewMatchTimerConfig creates an instance of the cognito configuration.
func NewMatchTimerConfig() *MatchTimerConfig {
	c := new(MatchTimerConfig)
	c.self = c
	c.Load("match_timer.yaml")
	return c
}

func (c *MatchTimerConfig) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	if err == nil {
		if err = yaml.Unmarshal(bytes, &c); err != nil {
			return err
		}
	}

	return nil
}

func (c *MatchTimerConfig) GetDTO() *game.MatchTimer {
	return &game.MatchTimer{
		Duration:     c.Duration,
		StoryMode:    c.StoryMode,
		ArenaMode:    c.ArenaMode,
		FriendlyMode: c.FriendlyMode,
		QuickMode:    c.QuickMode,
	}
}
