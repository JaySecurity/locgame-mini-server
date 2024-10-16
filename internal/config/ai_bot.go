package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// AiBot stores ai bot.
type AiBot struct {
	BaseConfig `yaml:"BaseConfig"`

	ProfileID string `yaml:"Profile"`
	DeckID    string `yaml:"Deck"`
}

// NewAiBot creates an instance of the ai bot configuration.
func NewAiBot() *AiBot {
	c := new(AiBot)
	c.self = c
	return c
}

func (c *AiBot) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bytes, &c)
}
