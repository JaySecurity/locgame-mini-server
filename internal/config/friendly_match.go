package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"locgame-mini-server/pkg/dto/friends"
)

func init() {
	Register(NewFriendlyMatch)
}

// FriendlyMatch stores friendly match configuration.
type FriendlyMatch struct {
	BaseConfig

	Data *friends.FriendlyMatchConfig
}

// NewFriendlyMatch creates an instance of the friendly match configuration.
func NewFriendlyMatch() *friends.FriendlyMatchConfig {
	c := new(FriendlyMatch)
	c.self = c
	c.Load("friendly_match.yaml")
	return c.Data
}

func (c *FriendlyMatch) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)

	if err == nil {
		err = yaml.Unmarshal(bytes, &c.Data)
	}

	return err
}
