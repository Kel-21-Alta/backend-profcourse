package quizs

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/quizs"
)

type QuizsRepository struct {
	Conn *gorm.DB
}

func (q QuizsRepository) CreateQuiz(ctx context.Context, domain *quizs.Domain) (quizs.Domain, error) {
	var rec = FromDomain(domain)
	err := q.Conn.Create(rec).Error

	if err != nil {
		return quizs.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) quizs.Repository {
	return &QuizsRepository{Conn: conn}
}
