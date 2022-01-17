package requests

import (
	"profcourse/business/quizs"
)

type UpdateQuizRequest struct {
	ID         string   `json:"id"`
	Pertanyaan string   `json:"pertanyaan"`
	Pilihan    []string `json:"pilihan"`
	Jawaban    string   `json:"jawaban"`
	ModulId    string   `json:"modul_id"`
}

func (receiver UpdateQuizRequest) ToDomain() *quizs.Domain {
	return &quizs.Domain{
		ID:         receiver.ID,
		Pilihan:    receiver.Pilihan,
		Pertanyaan: receiver.Pertanyaan,
		Jawaban:    receiver.Jawaban,
		ModulId:    receiver.ModulId,
	}
}
