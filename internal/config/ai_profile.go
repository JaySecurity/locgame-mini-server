package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// AiProfile stores ai strategy configuration.
type AiProfile struct {
	BaseConfig `yaml:"BaseConfig"`

	// Default target destroy weight
	DefaultDestroyWeight float32 `yaml:"DefaultDestroyWeight"`

	// Inflicted damage weight multiplier
	DamageWeightMultiplier float32 `yaml:"DamageWeightMultiplier"`

	// Serial targets destroy bonus
	SerialDestroyWeight float32 `yaml:"SerialDestroyWeight"`

	// Penalty multiplier for over damage on destroy
	OverDamagePenaltyMultiplier float32 `yaml:"OverDamagePenaltyMultiplier"`

	// Penalty multiplier for attack through strong edge
	StrongEdgeAttackPenaltyMultiplier float32 `yaml:"StrongEdgeAttackPenaltyMultiplier"`

	// Penalty for move that have possibility to attack, but don't attacks
	NonAttackingMovePenalty float32 `yaml:"NonAttackingMovePenalty"`

	// Penalty for move that loses
	LoseMovePenalty float32 `yaml:"LoseMovePenalty"`

	// Minimal edge defence weight multiplier
	MinimalDefenceWeight float32 `yaml:"MinimalDefenceWeight"`

	// Current HP weight multiplier
	CurrentHPWeight float32 `yaml:"CurrentHPWeight"`

	// Number of best moves to choose from all possible moves
	BestMovesCounter int32 `yaml:"BestMovesCounter"`

	// Ignore negative weight moves
	IgnoreNegativeWeightMoves bool `yaml:"IgnoreNegativeWeightMoves"`

	// Start percentile for choosing move from all possible moves
	ChooseStartIndex int32 `yaml:"ChooseStartIndex"`
}

// NewAiProfile creates an instance of the ai profile configuration.
func NewAiProfile() *AiProfile {
	c := new(AiProfile)
	c.self = c
	return c
}

func (c *AiProfile) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bytes, &c)
}
