package updateQuiz

import (
	"profcourse/business/quizs"
	"time"
)

type UpdateQuizResponse struct {
	ID         string    `json:"id"`
	Pertanyaan string    `json:"pertanyaan"`
	Pilihan    []string  `json:"pilihan"`
	Jawaban    string    `json:"jawaban"`
	ModulId    string    `json:"modul_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomain(domain quizs.Domain) *UpdateQuizResponse {
	return &UpdateQuizResponse{
		ID:         domain.ID,
		Pertanyaan: domain.Pertanyaan,
		Pilihan:    domain.Pilihan,
		Jawaban:    domain.Jawaban,
		ModulId:    domain.ModulId,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
