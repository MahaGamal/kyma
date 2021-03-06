// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import storage "github.com/kyma-project/kyma/components/ui-api-layer/internal/domain/content/storage"

// odataSpecGetter is an autogenerated failing mock type for the odataSpecGetter type
type odataSpecGetter struct {
	err error
}

// NewOdataSpecGetter creates a new odataSpecGetter type instance
func NewOdataSpecGetter(err error) *odataSpecGetter {
	return &odataSpecGetter{err: err}
}

// Find provides a failing mock function with given fields: kind, id
func (_m *odataSpecGetter) Find(kind string, id string) (*storage.ODataSpec, error) {
	var r0 *storage.ODataSpec
	var r1 error
	r1 = _m.err

	return r0, r1
}
