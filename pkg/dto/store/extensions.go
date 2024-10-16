package store

import "gopkg.in/yaml.v3"

//goland:noinspection GoMixedReceiverTypes
func (x *ProductType) MarshalYAML() (interface{}, error) {
	return ProductType_name[int32(*x)], nil
}

//goland:noinspection GoMixedReceiverTypes
func (x *ProductType) UnmarshalYAML(value *yaml.Node) error {
	*x = ProductType(ProductType_value[value.Value])
	return nil
}
