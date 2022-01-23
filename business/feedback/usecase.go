package feedback

import (
	"golang.org/x/net/context"
	controller "profcourse/controllers"
	"time"
)

type FeedbackUsecase struct {
	FeedbackRepository Repository
	ContextTimeOut     time.Duration
}

func (f FeedbackUsecase) CreateFeedback(ctx context.Context, domain *Domain) (Domain, error) {

	if domain.CourseId == "" {
		return Domain{}, controller.EMPTY_COURSE
	}

	if domain.Rating == 0.0 {
		return Domain{}, controller.RATING_EMPTY
	}

	if domain.Review == "" {
		return Domain{}, controller.REVIEW_EMPTY
	}

	if domain.UserId == "" {
		return Domain{}, controller.ID_EMPTY
	}

	result, err := f.FeedbackRepository.CreateFeedback(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func NewFeedbackUsecase(repo Repository, timeout time.Duration) Usecase {
	return &FeedbackUsecase{
		FeedbackRepository: repo,
		ContextTimeOut:     timeout,
	}
}
