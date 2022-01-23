package request_users

import (
	"context"
	controller "profcourse/controllers"
	"time"
)

type RequestUserUsecase struct {
	RequestUsercaseRepository Repository
	ContextTimeOut            time.Duration
}

func (r *RequestUserUsecase) CreateRequest(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Request == "" {
		return Domain{}, controller.REQUEST_EMPTY
	}
	if domain.UserId == "" {
		return Domain{}, controller.ID_EMPTY
	}
	// TODO: Pengecekan category ID
	result, err := r.RequestUsercaseRepository.CreateRequest(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func NewRequestUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &RequestUserUsecase{RequestUsercaseRepository: repo, ContextTimeOut: timeout}
}
