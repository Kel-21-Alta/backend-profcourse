package getOneQuiz

import (
	"profcourse/business/quizs"
	"time"
)

type GetOneQuizResponse struct {
	ID         string    `json:"id"`
	Pertanyaan string    `json:"pertanyaan"`
	Pilihan    []string  `json:"pilihan"`
	Jawaban    string    `json:"jawaban"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomain(domain quizs.Domain) *GetOneQuizResponse {
	return &GetOneQuizResponse{
		ID:         domain.ID,
		Pertanyaan: domain.Pertanyaan,
		Pilihan:    domain.Pilihan,
		Jawaban:    domain.Jawaban,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
