// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractMinProximity(opts ...interface{}) *opt.MinProximityOption {
	for _, o := range opts {
		if v, ok := o.(*opt.MinProximityOption); ok {
			return v
		}
	}
	return nil
}