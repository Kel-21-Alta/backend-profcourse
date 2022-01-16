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
}

type Usecase interface {
	CreateQuiz(ctx context.Context, domain *Domain) (Domain, error)
	ValidasiQuiz(ctx context.Context, domain *Domain) (*Domain, error)
}

type Repository interface {
	CreateQuiz(ctx context.Context, domain *Domain) (Domain, error)
}
