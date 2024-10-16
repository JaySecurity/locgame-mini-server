package arena

import "gopkg.in/yaml.v3"

//goland:noinspection GoMixedReceiverTypes
func (x *LeagueType) MarshalYAML() (interface{}, error) {
	return LeagueType_name[int32(*x)], nil
}

//goland:noinspection GoMixedReceiverTypes
func (x *LeagueType) UnmarshalYAML(value *yaml.Node) error {
	*x = LeagueType(LeagueType_value[value.Value])
	return nil
}
