package moduls

import (
	"context"
	controller "profcourse/controllers"
)

type modulUsecase struct {
	ModulRepository Repository
}

<<<<<<< HEAD
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
=======
func (m *modulUsecase) CreateModul(ctx context.Context, domain *Domain) (*Domain, error) {
>>>>>>> parent of f53e0f5... selesai membuat enpoint create modul
	if domain.Title == "" {
		return &Domain{}, controller.TITLE_EMPTY
	}
	if domain.Order <= 0 {
		return &Domain{}, controller.ORDER_MODUL_EMPTY
	}
	if domain.CourseId == "" {
		return &Domain{}, controller.EMPTY_COURSE
	}

	// TODO: cek apakah user adalah si pembuat course

	modul, err := m.ModulRepository.CreateModul(ctx, domain)

	if err != nil {
		return &Domain{}, err
	}

	return modul, nil
}

func NewModulUsecase(repository Repository) Usecase {
	return &modulUsecase{ModulRepository: repository}
}
