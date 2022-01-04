package summary

import (
	"golang.org/x/net/context"
	"profcourse/business/courses"
	"profcourse/business/users"
	"time"
)

type summaryUsecase struct {
	ContextTimeOut time.Duration
	CourseUsecase  courses.Usecase
	UserUsecase    users.Usecase
}

func (s *summaryUsecase) GetAllSummary(ctx context.Context) (*Domain, error) {

	courseDomain, err := s.CourseUsecase.GetCountCourse(ctx)
	if err != nil {
		return &Domain{}, err
	}

	userDomain, err := s.UserUsecase.GetCountUser(ctx)

	if err != nil {
		return &Domain{}, err
	}

	return &Domain{CountCourse: courseDomain.CountCourse, CountUser: userDomain.CountUser}, nil
}

func NewSummaryUsecase(timeout time.Duration, course courses.Usecase, user users.Usecase) Usecase {
	return &summaryUsecase{
		CourseUsecase:  course,
		ContextTimeOut: timeout,
		UserUsecase:    user,
	}
}
