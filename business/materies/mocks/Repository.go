// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	materies "profcourse/business/materies"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateMateri provides a mock function with given fields: ctx, domain
func (_m *Repository) CreateMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 materies.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *materies.Domain) materies.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(materies.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *materies.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMateri provides a mock function with given fields: ctx, domain
func (_m *Repository) DeleteMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 materies.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *materies.Domain) materies.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(materies.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *materies.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllMateri provides a mock function with given fields: ctx, domain
func (_m *Repository) GetAllMateri(ctx context.Context, domain *materies.Domain) (materies.AllMateriModul, error) {
	ret := _m.Called(ctx, domain)

	var r0 materies.AllMateriModul
	if rf, ok := ret.Get(0).(func(context.Context, *materies.Domain) materies.AllMateriModul); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(materies.AllMateriModul)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *materies.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCountMateriCourse provides a mock function with given fields: ctx, domain
func (_m *Repository) GetCountMateriCourse(ctx context.Context, domain *materies.Domain) (int, error) {
	ret := _m.Called(ctx, domain)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *materies.Domain) int); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *materies.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCountMateriFinish provides a mock function with given fields: ctx, domain
func (_m *Repository) GetCountMateriFinish(ctx context.Context, domain *materies.Domain) (int, error) {
	ret := _m.Called(ctx, domain)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *materies.Domain) int); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *materies.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOnemateri provides a mock function with given fields: ctx, domain
func (_m *Repository) GetOnemateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 materies.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *materies.Domain) materies.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(materies.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *materies.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMateri provides a mock function with given fields: ctx, domain
func (_m *Repository) UpdateMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 materies.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *materies.Domain) materies.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(materies.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *materies.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProgressMateri provides a mock function with given fields: ctx, domain
func (_m *Repository) UpdateProgressMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 materies.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *materies.Domain) materies.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(materies.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *materies.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
