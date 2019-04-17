// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

type ReferersOption struct {
	value []string
}

func Referers(v ...string) *ReferersOption {
	return &ReferersOption{v}
}

func (o *ReferersOption) Get() []string {
	if o == nil {
		return []string{}
	}
	return o.value
}

func (o ReferersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ReferersOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *ReferersOption) Equal(o2 *ReferersOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

func ReferersEqual(o1, o2 *ReferersOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
