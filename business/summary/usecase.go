package summary

import (
	"golang.org/x/net/context"
	"time"
)

type summaryUsecase struct {
	ContextTimeOut time.Duration
	SummaryRepository Repository
}

func (s summaryUsecase) GetSummary(ctx context.Context, domain *Domain) (*Domain, error) {
	summary, err := s.SummaryRepository.GetSummary(ctx, domain)
	if err != nil {
		return &Domain{}, err
	}
	return  summary, nil
}

func NewSummaryUsecase(r Repository, timeout time.Duration) Usecase {
	return &summaryUsecase{
		ContextTimeOut:    timeout,
		SummaryRepository: r,
	}
}