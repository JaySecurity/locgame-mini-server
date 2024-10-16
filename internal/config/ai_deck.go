package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// AiDeck stores ai deck.
type AiDeck struct {
	BaseConfig `yaml:"BaseConfig"`

	Cards []string `yaml:"Cards"`
}

// NewAiDeck creates an instance of the ai deck configuration.
func NewAiDeck() *AiDeck {
	c := new(AiDeck)
	c.self = c
	return c
}

func (c *AiDeck) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bytes, &c)
}
