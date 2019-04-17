// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type AutoGenerateObjectIDIfNotExistOption struct {
	value bool
}

func AutoGenerateObjectIDIfNotExist(v bool) *AutoGenerateObjectIDIfNotExistOption {
	return &AutoGenerateObjectIDIfNotExistOption{v}
}

func (o *AutoGenerateObjectIDIfNotExistOption) Get() bool {
	if o == nil {
		return false
	}
	return o.value
}

func (o AutoGenerateObjectIDIfNotExistOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AutoGenerateObjectIDIfNotExistOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *AutoGenerateObjectIDIfNotExistOption) Equal(o2 *AutoGenerateObjectIDIfNotExistOption) bool {
	if o2 == nil {
		return o.value == false
	}
	return o.value == o2.value
}

func AutoGenerateObjectIDIfNotExistEqual(o1, o2 *AutoGenerateObjectIDIfNotExistOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
