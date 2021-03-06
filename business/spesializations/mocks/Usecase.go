// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	spesializations "profcourse/business/spesializations"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// CreateSpasialization provides a mock function with given fields: ctx, domain
func (_m *Usecase) CreateSpasialization(ctx context.Context, domain *spesializations.Domain) (spesializations.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 spesializations.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *spesializations.Domain) spesializations.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(spesializations.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *spesializations.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllSpesializations provides a mock function with given fields: ctx, domain
func (_m *Usecase) GetAllSpesializations(ctx context.Context, domain *spesializations.Domain) ([]spesializations.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 []spesializations.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *spesializations.Domain) []spesializations.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]spesializations.Domain)
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

// GetCountSpesializations provides a mock function with given fields: ctx
func (_m *Usecase) GetCountSpesializations(ctx context.Context) (spesializations.Summary, error) {
	ret := _m.Called(ctx)

	var r0 spesializations.Summary
	if rf, ok := ret.Get(0).(func(context.Context) spesializations.Summary); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(spesializations.Summary)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneSpesialization provides a mock function with given fields: ctx, domain
func (_m *Usecase) GetOneSpesialization(ctx context.Context, domain *spesializations.Domain) (spesializations.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 spesializations.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *spesializations.Domain) spesializations.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(spesializations.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *spesializations.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
