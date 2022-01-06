package spesializations

import (
	"golang.org/x/net/context"
	controller "profcourse/controllers"
	"profcourse/helpers"
	"time"
)

type spesializationUsecase struct {
	SpesializationRepository Repository
	ContextTimeOut           time.Duration
}

func (s spesializationUsecase) GetAllSpesializations(ctx context.Context, domain *Domain) ([]Domain, error) {
	if domain.SortBy == "" {
		domain.SortBy = "asc"
	}

	if domain.SortBy == "dsc" {
		domain.SortBy = "desc"
	}

	if domain.Sort == "" {
		domain.Sort = "created_at"
	}

	// menvalidasi sort by yang diizinkan
	sortByAllow := []string{"asc", "desc"}
	if !helpers.CheckItemInSlice(sortByAllow, domain.SortBy) {
		return []Domain{}, controller.INVALID_PARAMS
	}

	// Menvalidasi sort yang diizinkan
	sortAllow := []string{"created_at", "title"} // TODO: disini kurang sort review dan sort popular
	if !helpers.CheckItemInSlice(sortAllow, domain.Sort) {
		return []Domain{}, controller.INVALID_PARAMS
	}

	result, err := s.SpesializationRepository.GetAllSpesializations(ctx, domain)

	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (s spesializationUsecase) CreateSpasialization(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.MakerRole != 1 {
		return Domain{}, controller.FORBIDDIN_USER
	}
	if domain.Title == "" {
		return Domain{}, controller.TITLE_EMPTY
	}
	if domain.Description == "" {
		return Domain{}, controller.DESC_EMPTY
	}
	if domain.ImageUrl == "" {
		return Domain{}, controller.IMAGE_EMPTY
	}
	if domain.Courses == nil || len(domain.Courses) < 1 {
		return Domain{}, controller.COURSES_SPESIALIZATION_EMPTY
	}
	result, err := s.SpesializationRepository.CreateSpasialization(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func NewSpesializationUsecase(r Repository, timeout time.Duration) Usecase {
	return &spesializationUsecase{
		SpesializationRepository: r,
		ContextTimeOut:           timeout,
	}
}
