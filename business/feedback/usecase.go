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

func (f FeedbackUsecase) DeleteFeedback(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.FEEDBACK_EMPTY
	}

	result, err := f.FeedbackRepository.DeleteFeedback(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, err
}

func (f FeedbackUsecase) GetAllFeedbackByCourse(ctx context.Context, domain *CourseReviews) (CourseReviews, error) {

	if domain.CourseId == "" {
		return CourseReviews{}, controller.EMPTY_COURSE
	}

	result, err := f.FeedbackRepository.GetAllFeedbackByCourse(ctx, domain)

	if err != nil {
		return CourseReviews{}, err
	}

	if len(result.Review) > 0 {
		ratingAll, err := f.FeedbackRepository.GetAvegareRatingCourse(ctx, domain)

		result.RatingAll = ratingAll.RatingAll

		if err != nil {
			return CourseReviews{}, err
		}
	}

	result.CourseId = domain.CourseId

	return result, nil
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
