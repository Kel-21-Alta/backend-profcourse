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

func (r *RequestUserUsecase) GetAllCategoryRequest(ctx context.Context) ([]Category, error) {
	result, err := r.RequestUsercaseRepository.GetAllCategoryRequest(ctx)

	if err != nil {
		return []Category{}, err
	}

	return result, nil
}

func (r *RequestUserUsecase) CreateRequest(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Request == "" {
		return Domain{}, controller.REQUEST_EMPTY
	}
	if domain.UserId == "" {
		return Domain{}, controller.ID_EMPTY
	}
	if domain.CategoryID == "" {
		return Domain{}, controller.CATEGORY_EMPTY
	}
	resultCreateReuqust, err := r.RequestUsercaseRepository.CreateRequest(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	result, err := r.RequestUsercaseRepository.GetOneRequest(ctx, &resultCreateReuqust)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func NewRequestUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &RequestUserUsecase{RequestUsercaseRepository: repo, ContextTimeOut: timeout}
}
