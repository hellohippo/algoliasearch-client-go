// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type GetRankingInfoOption struct {
	value bool
}

func GetRankingInfo(v bool) *GetRankingInfoOption {
	return &GetRankingInfoOption{v}
}

func (o *GetRankingInfoOption) Get() bool {
	if o == nil {
		return false
	}
	return o.value
}

func (o GetRankingInfoOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *GetRankingInfoOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *GetRankingInfoOption) Equal(o2 *GetRankingInfoOption) bool {
	if o2 == nil {
		return o.value == false
	}
	return o.value == o2.value
}

func GetRankingInfoEqual(o1, o2 *GetRankingInfoOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
