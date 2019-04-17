// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

type ExplainOption struct {
	value []string
}

func Explain(v ...string) *ExplainOption {
	return &ExplainOption{v}
}

func (o *ExplainOption) Get() []string {
	if o == nil {
		return []string{}
	}
	return o.value
}

func (o ExplainOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ExplainOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *ExplainOption) Equal(o2 *ExplainOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

func ExplainEqual(o1, o2 *ExplainOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
