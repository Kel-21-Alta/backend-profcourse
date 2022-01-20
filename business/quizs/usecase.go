package quizs

import (
	"golang.org/x/net/context"
	controller "profcourse/controllers"
	"time"
)

type QuizeUsecase struct {
	QuizRepository Repository
	ContextTimeOut time.Duration
}

func (q *QuizeUsecase) GetOneQuiz(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.ID == "" {
		return Domain{}, controller.ID_QUIZ_EMPTY
	}

	result, err := q.QuizRepository.GetOneQuiz(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (q *QuizeUsecase) GetAllQuizModul(ctx context.Context, domain *Domain) ([]Domain, error) {

	if domain.ModulId == "" {
		return []Domain{}, controller.EMPTY_MODUL_ID
	}

	result, err := q.QuizRepository.GetAllQuizModul(ctx, domain)

	if err != nil {
		return []Domain{}, err
	}

	return result, nil

}

func (q *QuizeUsecase) DeleteQuiz(ctx context.Context, id string) (string, error) {
	if id == "" {
		return "", controller.ID_QUIZ_EMPTY
	}

	result, err := q.QuizRepository.DeleteQuiz(ctx, id)

	if err != nil {
		return "", err
	}

	return result, nil
}

func (q *QuizeUsecase) UpdateQuiz(ctx context.Context, domain *Domain) (Domain, error) {

	if domain.ID == "" {
		return Domain{}, controller.ID_QUIZ_EMPTY
	}

	resultValidasi, err := q.ValidasiQuiz(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	result, err := q.QuizRepository.UpdateQuiz(ctx, resultValidasi)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (q *QuizeUsecase) ValidasiQuiz(ctx context.Context, domain *Domain) (*Domain, error) {
	if domain.ModulId == "" {
		return &Domain{}, controller.EMPTY_MODUL_ID
	}

	if domain.Jawaban == "" {
		return &Domain{}, controller.JAWABAN_QUIZ_EMPTY
	}

	if domain.Pertanyaan == "" {
		return &Domain{}, controller.PERTANYAAN_QUIZ_EMPTY
	}

	if domain.Pilihan == nil {
		return &Domain{}, controller.PILIHAN_QUIZ_EMPTY
	}

	if len(domain.Pilihan) < 2 {
		return &Domain{}, controller.PILIHAN_QUIZ_MINUS
	}

	return domain, nil
}

func (q *QuizeUsecase) CreateQuiz(ctx context.Context, domain *Domain) (Domain, error) {
	resultValidasi, err := q.ValidasiQuiz(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	result, err := q.QuizRepository.CreateQuiz(ctx, resultValidasi)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func NewQuizUsecase(repo Repository, timeout time.Duration) Usecase {
	return &QuizeUsecase{QuizRepository: repo, ContextTimeOut: timeout}
}
