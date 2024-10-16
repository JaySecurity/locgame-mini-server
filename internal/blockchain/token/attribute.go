package token

import (
	"encoding/json"
	"strconv"
)

type Attribute struct {
	DisplayType string         `json:"display_type"`
	TraitType   string         `json:"trait_type"`
	Value       AttributeValue `json:"value"`
}

type AttributeValue struct {
	isNumber bool

	intValue int32
	strValue string
}

func (v *AttributeValue) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		v.isNumber = true
		return json.Unmarshal(b, &v.intValue)
	} else {
		v.strValue = string(b[1 : len(b)-1])
	}
	return nil
}

func (v *AttributeValue) IsNumber() bool {
	return v.isNumber
}

func (v *AttributeValue) String() string {
	if v.isNumber {
		return strconv.Itoa(int(v.intValue))
	}
	return v.strValue
}

func (v *AttributeValue) Int() int32 {
	return v.intValue
}
