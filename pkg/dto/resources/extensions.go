package resources

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v3"
)

var resourcesByID map[int32]*ResourceData
var resourcesByKey map[string]*ResourceData

func SetResources(resources []*ResourceData) {
	resourcesByID = make(map[int32]*ResourceData)
	resourcesByKey = make(map[string]*ResourceData)

	for _, resource := range resources {
		resourcesByID[resource.ID] = resource
		resourcesByKey[resource.Key] = resource
	}
}

func (x *ResourceAdjustment) MarshalYAML() (interface{}, error) {
	return struct {
		Key      string `yaml:"Key"`
		Quantity int32  `yaml:"Quantity"`
	}{
		Key:      resourcesByID[x.ResourceID].Key,
		Quantity: x.Quantity,
	}, nil
}

func (x *ResourceAdjustment) UnmarshalYAML(value *yaml.Node) error {
	var data struct {
		Key      string
		Quantity int32
	}
	for i, node := range value.Content {
		switch node.Value {
		case "ResourceKey":
			fallthrough
		case "Key":
			data.Key = value.Content[i+1].Value
		case "Quantity":
			quantity, _ := strconv.Atoi(value.Content[i+1].Value)
			data.Quantity = int32(quantity)
		}
	}

	if _, ok := resourcesByKey[data.Key]; !ok {
		return fmt.Errorf("resource with key %s not found\n", data.Key)
	}

	*x = ResourceAdjustment{
		ResourceID: resourcesByKey[data.Key].ID,
		Quantity:   data.Quantity,
	}
	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (x CapacityType) MarshalYAML() (interface{}, error) {
	return CapacityType_name[int32(x)], nil
}

//goland:noinspection GoMixedReceiverTypes
func (x *CapacityType) UnmarshalYAML(value *yaml.Node) error {
	*x = CapacityType(CapacityType_value[value.Value])
	return nil
}
