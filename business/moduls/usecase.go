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

func (m *modulUsecase) CreateScoreModul(ctx context.Context, domain *ScoreUserModul) (ScoreUserModul, error) {
	if domain.ModulID == "" {
		return ScoreUserModul{}, controller.EMPTY_MODUL_ID
	}

	if domain.UserCourseId == "" {
		return ScoreUserModul{}, controller.EMPTY_COURSE
	}

	result, err := m.ModulRepository.CreateScoreModul(ctx, domain)

	if err != nil {
		return ScoreUserModul{}, err
	}

	return result, nil
}

func (m *modulUsecase) GetAllModulCourse(ctx context.Context, domain *Domain) ([]Domain, error) {
	if domain.CourseId == "" {
		return []Domain{}, controller.EMPTY_COURSE
	}

	result, err := m.ModulRepository.GetAllModulCourse(ctx, domain)

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (m *modulUsecase) DeleteModul(ctx context.Context, domain *Domain) (Message, error) {
	if domain.ID == "" {
		return "", controller.ID_EMPTY
	}

	// Mamastikan bahwa user yang akan mendelete modul adalah pemilik dari course atau admin
	modul, err := m.ModulRepository.GetOneModulWithCourse(ctx, domain)
	if err != nil {
		return "", err
	}

	if domain.RoleUser != 1 && domain.UserMakeModul != modul.UserMakeModul {
		return "", controller.FORBIDDIN_USER
	}

	massage, err := m.ModulRepository.DeleteModul(ctx, domain.ID)

	if err != nil {
		return "", err
	}

	return massage, nil
}

func (m *modulUsecase) UpdateModul(ctx context.Context, domain *Domain) (Domain, error) {

	if domain.Title == "" {
		return Domain{}, controller.TITLE_EMPTY
	}

	if domain.Order == 0 {
		return Domain{}, controller.ORDER_MODUL_EMPTY
	}

	if domain.CourseId == "" {
		return Domain{}, controller.EMPTY_COURSE
	}

	// cek yang melakukan pengubahan adalah admin atau pemilik dari course
	course, err := m.CourseUsecase.GetOneCourse(ctx, &courses.Domain{ID: domain.CourseId})

	if err != nil {
		return Domain{}, err
	}

	if (domain.RoleUser != 1) && (domain.UserMakeModul != course.TeacherId) {
		return Domain{}, controller.FORBIDDIN_USER
	}

	modul, err := m.ModulRepository.UpdateModul(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return modul, nil
}

func (m *modulUsecase) GetOneModul(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.EMPTY_MODUL_ID
	}

	result, err := m.ModulRepository.GetOneModul(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
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
