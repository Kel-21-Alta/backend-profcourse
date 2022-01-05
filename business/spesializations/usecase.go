package spesializations

import (
	"golang.org/x/net/context"
	controller "profcourse/controllers"
	"time"
)

type spesializationUsecase struct {
	SpesializationRepository Repository
	ContextTimeOut           time.Duration
}

func (s spesializationUsecase) CreateSpasialization(ctx context.Context, domain *Domain) (*Domain, error) {
	if domain.MakerRole != 1 {
		return &Domain{}, controller.FORBIDDIN_USER
	}
	if domain.Name == "" {
		return &Domain{}, controller.TITLE_EMPTY
	}
	if domain.Description == "" {
		return &Domain{}, controller.DESC_EMPTY
	}
	if domain.ImageUrl == "" {
		return &Domain{}, controller.IMAGE_EMPTY
	}
	if domain.Courses == nil || len(domain.Courses) < 1 {
		return &Domain{}, controller.COURSES_SPESIALIZATION_EMPTY
	}
	result, err := s.SpesializationRepository.CreateSpasialization(ctx, domain)

	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func NewSpesializationUsecase(r Repository, timeout time.Duration) Usecase {
	return &spesializationUsecase{
		SpesializationRepository: r,
		ContextTimeOut:           timeout,
	}
}
