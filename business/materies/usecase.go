package materies

import (
	"golang.org/x/net/context"
	controller "profcourse/controllers"
	"time"
)

type MateriesUsecase struct {
	MateriesRepository Repository
	ContextTimeout     time.Duration
}

func (u MateriesUsecase) GetAllMateri(ctx context.Context, domain *Domain) (AllMateriModul, error) {
	if domain.ModulId == "" {
		return AllMateriModul{}, controller.EMPTY_MODUL_ID
	}
	if domain.User.ID == "" {
		return AllMateriModul{}, controller.ID_EMPTY
	}

	result, err := u.MateriesRepository.GetAllMateri(ctx, domain)

	if err != nil {
		return AllMateriModul{}, err
	}

	return result, nil
}

func (u MateriesUsecase) GetOneMateri(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.ID_MATERI_EMPTY
	}

	// TODO: Validasi apakah user terdaftar di course materi ini

	result, err := u.MateriesRepository.GetOnemateri(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (u MateriesUsecase) DeleteMateri(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.ID_MATERI_EMPTY
	}

	result, err := u.MateriesRepository.DeleteMateri(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (u MateriesUsecase) ValidasiMateri(ctx context.Context, domain *Domain) (*Domain, error) {

	if domain.ModulId == "" {
		return &Domain{}, controller.EMPTY_MODUL_ID
	}
	// TODO: Validasi apakah modul id ada
	// TODO: Validasi apakah user yang sedang login adalah pemilik course atau admin
	if domain.Title == "" {
		return &Domain{}, controller.TITLE_EMPTY
	}
	if domain.UrlMateri == "" {
		return &Domain{}, controller.EMPTY_FILE_MATERI
	}
	if domain.Order == 0 {
		return &Domain{}, controller.ORDER_MATERI_EMPTY
	}
	if domain.Type == 0 {
		return &Domain{}, controller.TYPE_MATERI_EMPTY
	}
	if domain.Type < 1 || domain.Type > 2 {
		return &Domain{}, controller.TYPE_MATERI_WRONG
	}

	return domain, nil
}

func (u MateriesUsecase) UpdateMateri(ctx context.Context, domain *Domain) (Domain, error) {

	if domain.ID == "" {
		return Domain{}, controller.ID_MATERI_EMPTY
	}

	resultValidasi, err := u.ValidasiMateri(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	materi, err := u.MateriesRepository.UpdateMateri(ctx, resultValidasi)

	if err != nil {
		return Domain{}, err
	}

	return materi, nil
}

func (u MateriesUsecase) CreateMateri(ctx context.Context, domain *Domain) (Domain, error) {

	resultValidasi, err := u.ValidasiMateri(ctx, domain)

	materi, err := u.MateriesRepository.CreateMateri(ctx, resultValidasi)

	if err != nil {
		return Domain{}, err
	}

	return materi, nil
}

func NewMateriesUsecase(repo Repository, timeout time.Duration) Usecase {
	return &MateriesUsecase{MateriesRepository: repo}
}
