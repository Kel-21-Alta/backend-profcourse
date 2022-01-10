// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	courses "profcourse/business/courses"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateCourse provides a mock function with given fields: ctx, domain
func (_m *Repository) CreateCourse(ctx context.Context, domain *courses.Domain) (*courses.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 *courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *courses.Domain) *courses.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*courses.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *courses.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCourses provides a mock function with given fields: ctx, domain
func (_m *Repository) GetAllCourses(ctx context.Context, domain *courses.Domain) (*[]courses.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 *[]courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *courses.Domain) *[]courses.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]courses.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *courses.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCountCourse provides a mock function with given fields: ctx
func (_m *Repository) GetCountCourse(ctx context.Context) (*courses.Summary, error) {
	ret := _m.Called(ctx)

	var r0 *courses.Summary
	if rf, ok := ret.Get(0).(func(context.Context) *courses.Summary); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*courses.Summary)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneCourse provides a mock function with given fields: ctx, domain
func (_m *Repository) GetOneCourse(ctx context.Context, domain *courses.Domain) (*courses.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 *courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *courses.Domain) *courses.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*courses.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *courses.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCourse provides a mock function with given fields: ctx, domain
func (_m *Repository) UpdateCourse(ctx context.Context, domain *courses.Domain) (courses.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *courses.Domain) courses.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(courses.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *courses.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
