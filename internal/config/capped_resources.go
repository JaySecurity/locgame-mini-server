package config

import (
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v3"
	"locgame-mini-server/pkg/dto/resources"
)

func init() {
	Register(NewCappedResources)
}

// CappedResources stores capped resources configuration.
type CappedResources struct {
	BaseConfig

	Resources map[int32]*resources.CappedResource

	capacitiesOrdering map[int32][]int32 // Key - ResourceID, Value - Global Levels
}

// NewCappedResources creates an instance of capped resources' configuration.
func NewCappedResources() *CappedResources {
	c := new(CappedResources)
	c.self = c
	c.Load("capped_resources.yaml")
	return c
}

func (c *CappedResources) Unmarshal() error {
	bytes, err := ioutil.ReadFile(c.filePath)
	if err == nil {
		c.capacitiesOrdering = make(map[int32][]int32)
		var data []*resources.CappedResource
		err = yaml.Unmarshal(bytes, &data)
		if err == nil {
			c.Resources = make(map[int32]*resources.CappedResource)
			for _, resource := range data {
				c.Resources[resource.ResourceID] = resource

				var levels []int32

				for level := range resource.Capacities {
					levels = append(levels, level)
				}

				sort.Slice(levels, func(i, j int) bool { return levels[i] < levels[j] })

				c.capacitiesOrdering[resource.ResourceID] = levels
			}
		}
	}

	return err
}

func (c *CappedResources) HasHardCapacity(resourceID int32) bool {
	_, exists := c.Resources[resourceID]
	return exists && c.Resources[resourceID].CapacityType == resources.CapacityType_HardCapacity
}

func (c *CappedResources) GetCapacity(resourceID int32, level int32) int32 {
	var result int32 = 999_999_999

	if data, exists := c.Resources[resourceID]; exists {
		for _, capacityLevel := range c.capacitiesOrdering[resourceID] {
			if capacityLevel <= level {
				result = data.Capacities[capacityLevel]
			}
		}
	}

	return result
}
