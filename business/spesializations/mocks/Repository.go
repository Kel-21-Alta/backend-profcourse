// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	spesializations "profcourse/business/spesializations"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateSpasialization provides a mock function with given fields: ctx, domain
func (_m *Repository) CreateSpasialization(ctx context.Context, domain *spesializations.Domain) (*spesializations.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 *spesializations.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *spesializations.Domain) *spesializations.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spesializations.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *spesializations.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
