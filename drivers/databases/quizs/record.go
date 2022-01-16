package quizs

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"profcourse/business/quizs"
	"time"
)

type Quiz struct {
	ID         string `gorm:"primaryKey;unique"`
	Pertanyaan string `gorm:"not null"`
	Jawaban    string `gorm:"not null"`
	ModulId    string `gorm:"not null;size:191"`

	Pilihans []PilihanQuiz `gorm:"foreignKey:QuizId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type PilihanQuiz struct {
	ID      string `gorm:"primaryKey;unique"`
	Pilihan string `gorm:"not null"`
	QuizId  string `gorm:"not null;size:191"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (c *Quiz) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now().Local()
	return nil
}

func (c *Quiz) BeforeUpdate(db *gorm.DB) error {
	c.UpdatedAt = time.Now().Local()
	return nil
}

func (c *PilihanQuiz) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now().Local()
	return nil
}

func (c *PilihanQuiz) BeforeUpdate(db *gorm.DB) error {
	c.UpdatedAt = time.Now().Local()
	return nil
}

func (q Quiz) ToDomain() quizs.Domain {
	var pilihans []string
	for _, pilihan := range q.Pilihans {
		pilihans = append(pilihans, pilihan.Pilihan)
	}
	return quizs.Domain{
		ID:         q.ID,
		Pilihan:    pilihans,
		Pertanyaan: q.Pertanyaan,
		Jawaban:    q.Jawaban,
		ModulId:    q.ModulId,
		CreatedAt:  q.CreatedAt,
		UpdatedAt:  q.UpdatedAt,
	}

}

func FromDomain(domain *quizs.Domain) *Quiz {
	var listPilihan []PilihanQuiz
	for _, pilihan := range domain.Pilihan {
		listPilihan = append(listPilihan, PilihanQuiz{
			Pilihan: pilihan,
		})
	}
	return &Quiz{
		ID:         domain.ID,
		Pertanyaan: domain.Pertanyaan,
		Pilihans:   listPilihan,
		Jawaban:    domain.Jawaban,
		ModulId:    domain.ModulId,
	}
}
