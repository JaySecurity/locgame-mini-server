package config

import (
	"io/ioutil"
	"locgame-mini-server/pkg/dto/game"
	"locgame-mini-server/pkg/dto/resources"

	"gopkg.in/yaml.v3"
)

func init() {
	Register(NewExtraRewardsConfig)
}

// ExtraRewardsConfig stores extra rewards configuration.
type ExtraRewardsConfig struct {
	BaseConfig

	ComboAttacks map[int32]*resources.ResourceAdjustment `yaml:"ComboAttacks"`
	//get reward by visual rarity of the card
	CenterStage map[int32]*resources.ResourceAdjustment `yaml:"CenterStage"`
	//get reward by visual rarity of the cards
	CardsInField   map[int32]*resources.ResourceAdjustment `yaml:"CardsInField"`
	StoryMode      bool                                    `yaml:"StoryMode"`
	ArenaMode      bool                                    `yaml:"ArenaMode"`
	FriendlyMode   bool                                    `yaml:"FriendlyMode"`
	QuickMatchMode bool                                    `yaml:"QuickMatchMode"`
}

// NewExtraRewardsConfig creates an instance of the extra rewards configuration.
func NewExtraRewardsConfig() *ExtraRewardsConfig {
	c := new(ExtraRewardsConfig)
	c.self = c
	c.Load("extra_rewards.yaml")
	return c
}

func (c *ExtraRewardsConfig) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	if err == nil {
		err = yaml.Unmarshal(bytes, &c)
	}

	return err
}

func (c *ExtraRewardsConfig) GetDTO() *game.ExtraRewards {
	return &game.ExtraRewards{
		ComboAttacks:   c.ComboAttacks,
		CenterStage:    c.CenterStage,
		CardsInField:   c.CardsInField,
		StoryMode:      c.StoryMode,
		ArenaMode:      c.ArenaMode,
		FriendlyMode:   c.FriendlyMode,
		QuickMatchMode: c.QuickMatchMode,
	}
}
