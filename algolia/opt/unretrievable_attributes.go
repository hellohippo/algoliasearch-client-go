// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

type UnretrievableAttributesOption struct {
	value []string
}

func UnretrievableAttributes(v ...string) *UnretrievableAttributesOption {
	return &UnretrievableAttributesOption{v}
}

func (o *UnretrievableAttributesOption) Get() []string {
	if o == nil {
		return []string{}
	}
	return o.value
}

func (o UnretrievableAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *UnretrievableAttributesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *UnretrievableAttributesOption) Equal(o2 *UnretrievableAttributesOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

func UnretrievableAttributesEqual(o1, o2 *UnretrievableAttributesOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
