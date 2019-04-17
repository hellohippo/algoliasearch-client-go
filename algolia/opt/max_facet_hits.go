// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type MaxFacetHitsOption struct {
	value int
}

func MaxFacetHits(v int) *MaxFacetHitsOption {
	return &MaxFacetHitsOption{v}
}

func (o *MaxFacetHitsOption) Get() int {
	if o == nil {
		return 10
	}
	return o.value
}

func (o MaxFacetHitsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *MaxFacetHitsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 10
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *MaxFacetHitsOption) Equal(o2 *MaxFacetHitsOption) bool {
	if o2 == nil {
		return o.value == 10
	}
	return o.value == o2.value
}

func MaxFacetHitsEqual(o1, o2 *MaxFacetHitsOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
