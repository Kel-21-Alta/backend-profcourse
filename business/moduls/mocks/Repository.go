// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	moduls "profcourse/business/moduls"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateModul provides a mock function with given fields: ctx, domain
func (_m *Repository) CreateModul(ctx context.Context, domain *moduls.Domain) (moduls.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 moduls.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *moduls.Domain) moduls.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(moduls.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *moduls.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneModul provides a mock function with given fields: ctx, domain
func (_m *Repository) GetOneModul(ctx context.Context, domain *moduls.Domain) (moduls.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 moduls.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *moduls.Domain) moduls.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(moduls.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *moduls.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
