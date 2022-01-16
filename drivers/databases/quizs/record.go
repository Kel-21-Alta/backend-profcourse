package quizs

import (
	"gorm.io/gorm"
	"profcourse/business/quizs"
	"time"
)

type Quiz struct {
	ID         string   `gorm:"primaryKey;unique"`
	Pertanyaan string   `gorm:"not null"`
	Pilihan    []string `gorm:"not null"`
	Jawaban    string   `gorm:"not null"`
	ModulId    string   `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (q Quiz) ToDomain() quizs.Domain {
	return quizs.Domain{
		ID:         q.ID,
		Pilihan:    q.Pilihan,
		Pertanyaan: q.Pertanyaan,
		Jawaban:    q.Jawaban,
		ModulId:    q.ModulId,
		CreatedAt:  q.CreatedAt,
		UpdatedAt:  q.UpdatedAt,
	}

}

func FromDomain(domain *quizs.Domain) *Quiz {
	return &Quiz{
		ID:         domain.ID,
		Pertanyaan: domain.Pertanyaan,
		Pilihan:    domain.Pilihan,
		Jawaban:    domain.Jawaban,
		ModulId:    domain.ModulId,
	}
}
