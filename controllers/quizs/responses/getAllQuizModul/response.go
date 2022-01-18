package getAllQuizModul

import (
	"profcourse/business/quizs"
	"time"
)

type GetAllQuizModulResponse struct {
	ID         string    `json:"id"`
	Pertanyaan string    `json:"pertanyaan"`
	Pilihan    []string  `json:"pilihan"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromListDomain(domains []quizs.Domain) []GetAllQuizModulResponse {
	var listResponse []GetAllQuizModulResponse

	for _, quiz := range domains {
		listResponse = append(listResponse, GetAllQuizModulResponse{
			ID:         quiz.ID,
			Pertanyaan: quiz.Pertanyaan,
			Pilihan:    quiz.Pilihan,
			CreatedAt:  quiz.CreatedAt,
			UpdatedAt:  quiz.UpdatedAt,
		})
	}

	return listResponse
}
