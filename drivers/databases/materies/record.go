package materies

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type TYPE string

const (
	Text  TYPE = "text"
	Video TYPE = "video"
)

type Materi struct {
	ID        string `gorm:"primaryKey;unique"`
	Title     string `gorm:"not null"`
	ModulID   string `gorm:"not null;size:191"`
	Order     int8   `gorm:"not null"`
	Type      TYPE   `gorm:"not null"`
	UrlMateri string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *Materi) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewV4().String()
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *Materi) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
