package requests

import (
	"profcourse/business/quizs"
)

type CreateQuizRequest struct {
	Pertanyaan string   `json:"pertanyaan"`
	Pilihan    []string `json:"pilihan"`
	Jawaban    string   `json:"jawaban"`
	ModulId    string   `json:"modul_id"`
}

func (r *CreateQuizRequest) ToDomain() *quizs.Domain {
	return &quizs.Domain{
		Pilihan:    r.Pilihan,
		Pertanyaan: r.Pertanyaan,
		Jawaban:    r.Jawaban,
		ModulId:    r.ModulId,
	}
}
