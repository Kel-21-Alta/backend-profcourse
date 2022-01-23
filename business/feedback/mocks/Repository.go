// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	feedback "profcourse/business/feedback"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateFeedback provides a mock function with given fields: ctx, domain
func (_m *Repository) CreateFeedback(ctx context.Context, domain *feedback.Domain) (feedback.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 feedback.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *feedback.Domain) feedback.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(feedback.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *feedback.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllFeedbackByCourse provides a mock function with given fields: ctx, domain
func (_m *Repository) GetAllFeedbackByCourse(ctx context.Context, domain *feedback.CourseReviews) (feedback.CourseReviews, error) {
	ret := _m.Called(ctx, domain)

	var r0 feedback.CourseReviews
	if rf, ok := ret.Get(0).(func(context.Context, *feedback.CourseReviews) feedback.CourseReviews); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(feedback.CourseReviews)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *feedback.CourseReviews) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAvegareRatingCourse provides a mock function with given fields: ctx, domain
func (_m *Repository) GetAvegareRatingCourse(ctx context.Context, domain *feedback.CourseReviews) (feedback.CourseReviews, error) {
	ret := _m.Called(ctx, domain)

	var r0 feedback.CourseReviews
	if rf, ok := ret.Get(0).(func(context.Context, *feedback.CourseReviews) feedback.CourseReviews); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(feedback.CourseReviews)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *feedback.CourseReviews) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneFeedbackByUserAndCourse provides a mock function with given fields: ctx, domain
func (_m *Repository) GetOneFeedbackByUserAndCourse(ctx context.Context, domain *feedback.Domain) (feedback.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 feedback.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *feedback.Domain) feedback.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(feedback.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *feedback.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFeedback provides a mock function with given fields: ctx, domain
func (_m *Repository) UpdateFeedback(ctx context.Context, domain *feedback.Domain) (feedback.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 feedback.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *feedback.Domain) feedback.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(feedback.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *feedback.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
