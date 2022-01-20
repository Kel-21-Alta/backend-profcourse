package quizs

import (
	"golang.org/x/net/context"
	"profcourse/business/moduls"
	"profcourse/business/users_courses"
	controller "profcourse/controllers"
	"time"
)

type QuizeUsecase struct {
	QuizRepository Repository
	ContextTimeOut time.Duration
	ModulUsecase   moduls.Usecase
	UserCourseUsecase users_courses.Usecase
}

func (q *QuizeUsecase) CalculateScoreQuiz(ctx context.Context, domain []Domain, userId string) (Domain, error) {
	var result Domain

	for _, dom := range domain {
		if dom.ModulId == "" {
			return Domain{}, controller.EMPTY_MODUL_ID
		}
		if dom.ID == "" {
			return Domain{}, controller.ID_QUIZ_EMPTY
		}
		result.ModulId = dom.ModulId
	}

	allQuizModul, err := q.QuizRepository.GetAllQuizModul(ctx, &Domain{ModulId: result.ModulId})

	if err != nil {
		return Domain{}, err
	}

	skor := 0

	for _, quiz := range allQuizModul {
		for _, jawaban := range domain {
			if quiz.ID == jawaban.ID {
				if jawaban.Jawaban == "" {
					skor += 0
				} else if quiz.Jawaban == jawaban.Jawaban {
					skor += 2
				} else {
					skor -= 1
				}
			}
		}
	}

	modul, err := q.ModulUsecase.GetOneModul(ctx, &moduls.Domain{ID: result.ModulId})

	if err != nil {
		return Domain{}, err
	}

	userCourse, err := q.UserCourseUsecase.GetOneUserCourse(ctx, &users_courses.Domain{
		UserId:      userId,
		CourseId:    modul.CourseId,
	})

	if err != nil {
		return Domain{}, err
	}

	// Create score quiz modul
	resultInputScore, err := q.ModulUsecase.CreateScoreModul(ctx, &moduls.ScoreUserModul{
		Nilai:        skor,
		ModulID:      result.ModulId,
		UserCourseId: userCourse.ID,
	})

	if err != nil {
		return Domain{}, err
	}

	// Mengupdate score course
	userCourse.Score += skor
	_, err = q.UserCourseUsecase.UpdateScoreCourse(ctx, &userCourse)

	if err != nil {
		return Domain{}, err
	}
	result.Skor = skor
	result.ModulId = resultInputScore.ModulID
	result.ModulTitle = modul.Title

	return result, nil
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

func NewQuizUsecase(repo Repository, modul moduls.Usecase, userCourse users_courses.Usecase, timeout time.Duration) Usecase {
	return &QuizeUsecase{QuizRepository: repo, ContextTimeOut: timeout, ModulUsecase: modul, UserCourseUsecase: userCourse}
}

