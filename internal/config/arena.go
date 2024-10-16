package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"locgame-mini-server/pkg/dto/resources"
)

func init() {
	Register(NewArenaConfig)
}

// ArenaConfig stores arena configuration.
type ArenaConfig struct {
	BaseConfig

	Matchmaking struct {
		Range struct {
			Min int32 `yaml:"Min"`
			Max int32 `yaml:"Max"`
		} `yaml:"range"`
		IncreasedRange struct {
			Min int32 `yaml:"Min"`
			Max int32 `yaml:"Max"`
		} `yaml:"IncreasedRange"`
	} `yaml:"Matchmaking"`
	DefaultAvatar  string                        `yaml:"DefaultAvatar"`
	Reward         *resources.ResourceAdjustment `yaml:"Reward"`
	ReviveCost     *resources.ResourceAdjustment `yaml:"ReviveCost"`
	TicketCost     *resources.ResourceAdjustment `yaml:"TicketCost"`
	FreeTicketCost *resources.ResourceAdjustment `yaml:"FreeTicketCost"`
	AiProfiles     map[int32][]struct {
		Min     int32  `yaml:"Min"`
		Max     int32  `yaml:"Max"`
		Profile string `yaml:"Profile"`
	} `yaml:"AiProfiles"`
}

// NewArenaConfig creates an instance of the arena configuration.
func NewArenaConfig() *ArenaConfig {
	c := new(ArenaConfig)
	c.self = c
	c.Load("arena.yaml")
	return c
}

func (c *ArenaConfig) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	if err == nil {
		if err = yaml.Unmarshal(bytes, &c); err != nil {
			return err
		}
	}

	return nil
}
