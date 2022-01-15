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

func (u MateriesUsecase) CreateMateri(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.ModulId == "" {
		return Domain{}, controller.EMPTY_MODUL_ID
	}
	// TODO: Validasi apakah modul id ada
	// TODO: Validasi apakah user yang sedang login adalah pemilik course atau admin
	if domain.Title == "" {
		return Domain{}, controller.TITLE_EMPTY
	}
	if domain.UrlMateri == "" {
		return Domain{}, controller.EMPTY_FILE_MATERI
	}
	if domain.Order == 0 {
		return Domain{}, controller.ORDER_MATERI_EMPTY
	}
	if domain.Type == 0 {
		return Domain{}, controller.TYPE_MATERI_EMPTY
	}

	if domain.Type < 1 || domain.Type > 2 {
		return Domain{}, controller.TYPE_MATERI_WRONG
	}

	materi, err := u.MateriesRepository.CreateMateri(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return materi, nil
}

func NewMateriesUsecase(repo Repository, timeout time.Duration) Usecase {
	return &MateriesUsecase{MateriesRepository: repo}
}
