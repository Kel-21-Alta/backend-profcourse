// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	users_courses "profcourse/business/users_courses"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetEndRollCourseUserById provides a mock function with given fields: ctx, domain
func (_m *Repository) GetEndRollCourseUserById(ctx context.Context, domain *users_courses.Domain) (*users_courses.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 *users_courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *users_courses.Domain) *users_courses.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users_courses.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users_courses.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneUserCourse provides a mock function with given fields: ctx, domain
func (_m *Repository) GetOneUserCourse(ctx context.Context, domain *users_courses.Domain) (users_courses.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 users_courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *users_courses.Domain) users_courses.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(users_courses.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users_courses.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProgressCourse provides a mock function with given fields: ctx, domain
func (_m *Repository) UpdateProgressCourse(ctx context.Context, domain *users_courses.Domain) (users_courses.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 users_courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *users_courses.Domain) users_courses.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(users_courses.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users_courses.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRegisterCourse provides a mock function with given fields: ctx, domain
func (_m *Repository) UserRegisterCourse(ctx context.Context, domain *users_courses.Domain) (*users_courses.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 *users_courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *users_courses.Domain) *users_courses.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users_courses.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users_courses.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
