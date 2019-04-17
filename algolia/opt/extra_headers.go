// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

type ExtraHeadersOption struct {
	value map[string]string
}

func ExtraHeaders(v map[string]string) *ExtraHeadersOption {
	return &ExtraHeadersOption{v}
}

func (o *ExtraHeadersOption) Get() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.value
}

func (o ExtraHeadersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ExtraHeadersOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = map[string]string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *ExtraHeadersOption) Equal(o2 *ExtraHeadersOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, map[string]string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

func ExtraHeadersEqual(o1, o2 *ExtraHeadersOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
