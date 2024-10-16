package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"locgame-mini-server/pkg/dto/resources"
)

func init() {
	Register(NewResettableResources)
}

// ResettableResources stores resettable resources configuration.
type ResettableResources struct {
	BaseConfig

	*resources.ResettableResources
}

// NewResettableResources creates an instance of resettable resources' configuration.
func NewResettableResources() map[int32]*resources.ResettableResource {
	c := new(ResettableResources)
	c.self = c
	c.Load("resettable_resources.yaml")
	return c.Resources
}

func (c *ResettableResources) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)
	if err == nil {
		var data []*resources.ResettableResource
		err = yaml.Unmarshal(bytes, &data)
		if err == nil {
			c.ResettableResources = new(resources.ResettableResources)
			c.ResettableResources.Resources = make(map[int32]*resources.ResettableResource)
			for _, resource := range data {
				c.ResettableResources.Resources[resource.ResourceID] = resource
			}
		}
	}

	return err
}
