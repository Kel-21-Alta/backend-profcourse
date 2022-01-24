package quizs

import (
	"golang.org/x/net/context"
	"time"
)

type Domain struct {
	ID         string
	Pilihan    []string
	Pertanyaan string
	Jawaban    string
	ModulId    string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Skor       int
	ModulTitle string
}

type Usecase interface {
	CreateQuiz(ctx context.Context, domain *Domain) (Domain, error)
	ValidasiQuiz(ctx context.Context, domain *Domain) (*Domain, error)
	UpdateQuiz(ctx context.Context, domain *Domain) (Domain, error)
	DeleteQuiz(ctx context.Context, id string) (string, error)
	GetAllQuizModul(ctx context.Context, domain *Domain) ([]Domain, error)
	GetOneQuiz(ctx context.Context, domain *Domain) (Domain, error)
	CalculateScoreQuiz(ctx context.Context, domain []Domain, userId string) (Domain, error)
}

type Repository interface {
	CreateQuiz(ctx context.Context, domain *Domain) (Domain, error)
	UpdateQuiz(ctx context.Context, domain *Domain) (Domain, error)
	DeleteQuiz(ctx context.Context, id string) (string, error)
	GetAllQuizModul(ctx context.Context, domain *Domain) ([]Domain, error)
	GetOneQuiz(ctx context.Context, domain *Domain) (Domain, error)
}
