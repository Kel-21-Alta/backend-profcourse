package requests

import "profcourse/business/quizs"

type JawabanQuiz struct {
	QuizId  string `json:"quiz_id"`
	Jawaban string `json:"jawaban"`
}

type RequestJawabans struct {
	Jawaban []JawabanQuiz `json:"jawaban"`
	ModulId string `json:"modul_id"`
}

func (j *JawabanQuiz) ToDomain() quizs.Domain {
	return quizs.Domain{
		ID: j.QuizId,
		Jawaban: j.Jawaban,
	}
}

func ToListDomain(request RequestJawabans) []quizs.Domain {
	var listDomain []quizs.Domain

	for _, req := range request.Jawaban {
		domain := req.ToDomain()
		domain.ModulId = request.ModulId
		listDomain = append(listDomain,domain)
	}

	return listDomain
}