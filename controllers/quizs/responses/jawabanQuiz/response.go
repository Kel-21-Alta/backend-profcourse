package jawabanQuiz

import "profcourse/business/quizs"

type JawabanQuizResponse struct {
	Skor int `json:"skor"`
	ModulId    string `json:"modul_id"`
	ModulTitle string `json:"modul_title"`
}

func FromDomain(domain quizs.Domain) *JawabanQuizResponse  {
	return &JawabanQuizResponse{Skor: domain.Skor, ModulTitle: domain.ModulTitle, ModulId: domain.ModulId}
}
