// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestDisableTypoToleranceOnAttributes(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.DisableTypoToleranceOnAttributesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.DisableTypoToleranceOnAttributes(),
		},
		{
			opts:     []interface{}{opt.DisableTypoToleranceOnAttributes("value1")},
			expected: opt.DisableTypoToleranceOnAttributes("value1"),
		},
		{
			opts:     []interface{}{opt.DisableTypoToleranceOnAttributes("value1", "value2", "value3")},
			expected: opt.DisableTypoToleranceOnAttributes("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractDisableTypoToleranceOnAttributes(c.opts...)
			out opt.DisableTypoToleranceOnAttributesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}