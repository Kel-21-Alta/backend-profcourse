package moduls

import (
	"context"
	controller "profcourse/controllers"
)

type modulUsecase struct {
	ModulRepository Repository
}

func (m modulUsecase) CreateModul(ctx context.Context, domain *Domain) (*Domain, error) {
	if domain.Title == "" {
		return &Domain{}, controller.TITLE_EMPTY
	}
	if domain.Order <= 0 {
		return &Domain{}, controller.ORDER_MODUL_EMPTY
	}
	if domain.CourseId == "" {
		return &Domain{}, controller.EMPTY_COURSE
	}
	modul, err := m.ModulRepository.CreateModul(ctx, domain)

	if err != nil {
		return &Domain{}, err
	}

	return modul, nil
}

func NewModulRepository(repository Repository) Usecase {
	return &modulUsecase{ModulRepository: repository}
}
