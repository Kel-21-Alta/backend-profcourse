package moduls

import (
	"context"
	"profcourse/business/courses"
	controller "profcourse/controllers"
	"time"
)

type modulUsecase struct {
	ModulRepository Repository
	ContextTimeOut  time.Duration
	CourseUsecase   courses.Usecase
}

func (m *modulUsecase) CreateModul(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Title == "" {
		return Domain{}, controller.TITLE_EMPTY
	}
	if domain.Order <= 0 {
		return Domain{}, controller.ORDER_MODUL_EMPTY
	}
	if domain.CourseId == "" {
		return Domain{}, controller.EMPTY_COURSE
	}

	course, err := m.CourseUsecase.GetOneCourse(ctx, &courses.Domain{ID: domain.CourseId})

	if err != nil {
		return Domain{}, err
	}

	// Validasi hanya admin dan user yang membuat course yang dapat menambahkan modul
	if domain.RoleUser != 1 {
		if course.TeacherId != domain.UserMakeModul {
			return Domain{}, controller.FORBIDDIN_USER
		}
	}

	modul, err := m.ModulRepository.CreateModul(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return modul, nil
}

func NewModulUsecase(repository Repository, courseUsecase courses.Usecase, timeout time.Duration) Usecase {
	return &modulUsecase{ModulRepository: repository, CourseUsecase: courseUsecase, ContextTimeOut: timeout}
}
