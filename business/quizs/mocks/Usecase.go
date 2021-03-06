// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	quizs "profcourse/business/quizs"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// CalculateScoreQuiz provides a mock function with given fields: ctx, domain, userId
func (_m *Usecase) CalculateScoreQuiz(ctx context.Context, domain []quizs.Domain, userId string) (quizs.Domain, error) {
	ret := _m.Called(ctx, domain, userId)

	var r0 quizs.Domain
	if rf, ok := ret.Get(0).(func(context.Context, []quizs.Domain, string) quizs.Domain); ok {
		r0 = rf(ctx, domain, userId)
	} else {
		r0 = ret.Get(0).(quizs.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []quizs.Domain, string) error); ok {
		r1 = rf(ctx, domain, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateQuiz provides a mock function with given fields: ctx, domain
func (_m *Usecase) CreateQuiz(ctx context.Context, domain *quizs.Domain) (quizs.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 quizs.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *quizs.Domain) quizs.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(quizs.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *quizs.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteQuiz provides a mock function with given fields: ctx, id
func (_m *Usecase) DeleteQuiz(ctx context.Context, id string) (string, error) {
	ret := _m.Called(ctx, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllQuizModul provides a mock function with given fields: ctx, domain
func (_m *Usecase) GetAllQuizModul(ctx context.Context, domain *quizs.Domain) ([]quizs.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 []quizs.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *quizs.Domain) []quizs.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]quizs.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *quizs.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneQuiz provides a mock function with given fields: ctx, domain
func (_m *Usecase) GetOneQuiz(ctx context.Context, domain *quizs.Domain) (quizs.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 quizs.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *quizs.Domain) quizs.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(quizs.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *quizs.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateQuiz provides a mock function with given fields: ctx, domain
func (_m *Usecase) UpdateQuiz(ctx context.Context, domain *quizs.Domain) (quizs.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 quizs.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *quizs.Domain) quizs.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(quizs.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *quizs.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidasiQuiz provides a mock function with given fields: ctx, domain
func (_m *Usecase) ValidasiQuiz(ctx context.Context, domain *quizs.Domain) (*quizs.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 *quizs.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *quizs.Domain) *quizs.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*quizs.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *quizs.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
