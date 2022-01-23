// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	users_courses "profcourse/business/users_courses"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetOneUserCourse provides a mock function with given fields: ctx, domain
func (_m *Usecase) GetOneUserCourse(ctx context.Context, domain *users_courses.Domain) (users_courses.Domain, error) {
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

// GetUserCourseEndroll provides a mock function with given fields: ctx, domain
func (_m *Usecase) GetUserCourseEndroll(ctx context.Context, domain *users_courses.User) (users_courses.User, error) {
	ret := _m.Called(ctx, domain)

	var r0 users_courses.User
	if rf, ok := ret.Get(0).(func(context.Context, *users_courses.User) users_courses.User); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(users_courses.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users_courses.User) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProgressCourse provides a mock function with given fields: ctx, domain
func (_m *Usecase) UpdateProgressCourse(ctx context.Context, domain *users_courses.Domain) (users_courses.Domain, error) {
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

// UpdateScoreCourse provides a mock function with given fields: ctx, domain
func (_m *Usecase) UpdateScoreCourse(ctx context.Context, domain *users_courses.Domain) (users_courses.Domain, error) {
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
func (_m *Usecase) UserRegisterCourse(ctx context.Context, domain *users_courses.Domain) (*users_courses.Domain, error) {
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
