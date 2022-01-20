package quizs

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/quizs"
)

type QuizsRepository struct {
	Conn *gorm.DB
}

func (q QuizsRepository) GetOneQuiz(ctx context.Context, domain *quizs.Domain) (quizs.Domain, error) {
	var rec Quiz

	err := q.Conn.Preload("Pilihans").First(&rec, "id = ?", domain.ID).Error

	if err != nil {
		return quizs.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (q QuizsRepository) GetAllQuizModul(ctx context.Context, domain *quizs.Domain) ([]quizs.Domain, error) {
	var rec []Quiz

	err := q.Conn.Preload("Pilihans").Where("modul_id = ?", domain.ModulId).Find(&rec).Order("created_at desc").Error

	if err != nil {
		return []quizs.Domain{}, err
	}

	return ToListDomain(rec), nil
}

func (q QuizsRepository) DeleteQuiz(ctx context.Context, id string) (string, error) {
	var rec Quiz
	var rec2 PilihanQuiz

	if err := q.Conn.Where("quiz_id = ?", id).Delete(&rec2).Error; err != nil {
		return "", err
	}

	if err := q.Conn.Delete(&rec, "id = ?", id).Error; err != nil {
		return "", err
	}

	return id, nil
}

func (q QuizsRepository) UpdateQuiz(ctx context.Context, domain *quizs.Domain) (quizs.Domain, error) {
	var rec Quiz
	var err error
	err = q.Conn.First(&rec, "id = ?", domain.ID).Error
	if err != nil {
		return quizs.Domain{}, err
	}

	rec.Pertanyaan = domain.Pertanyaan
	rec.Jawaban = domain.Jawaban
	rec.ModulId = domain.ModulId

	//err = q.Conn.Model(&rec).Association("Pilihans").Clear()
	var pilihanQuiz PilihanQuiz

	err = q.Conn.Where("quiz_id = ?", domain.ID).Unscoped().Delete(&pilihanQuiz).Error

	if err != nil {
		return quizs.Domain{}, err
	}

	var listPilihan []PilihanQuiz

	for _, pilihan := range domain.Pilihan {
		listPilihan = append(listPilihan, PilihanQuiz{Pilihan: pilihan})
	}

	rec.Pilihans = listPilihan

	err = q.Conn.Save(&rec).Error

	if err != nil {
		return quizs.Domain{}, err
	}

	return rec.ToDomain(), nil
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
