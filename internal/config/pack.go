package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	storeDto "locgame-mini-server/pkg/dto/store"
)

// Pack stores pack configuration.
type Pack struct {
	BaseConfig
	*storeDto.Pack
}

// NewPack creates an instance of the pack configuration.
func NewPack() *Pack {
	c := new(Pack)
	c.self = c
	return c
}

func (c *Pack) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bytes, &c.Pack)
}
