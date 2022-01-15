package materies

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"profcourse/business/materies"
	"time"
)

type TYPE int8

const (
	Text  TYPE = 1
	Video TYPE = 2
)

type Materi struct {
	ID        string `gorm:"primaryKey;unique"`
	Title     string `gorm:"not null"`
	ModulID   string `gorm:"not null;size:191"`
	Order     int8   `gorm:"not null"`
	Type      TYPE   `gorm:"not null"`
	UrlMateri string `gorm:"not null"`

	MateriUserComplate []MateriUserComplate `gorm:"foreignKey:MateriId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type MateriUserComplate struct {
	ID          string `gorm:"primaryKey;unique"`
	MateriId    string `gorm:"not null;size:191;index:idx_unique1,unique"`
	UserId      string `gorm:"not null;size:191;index:idx_unique1,unique"`
	CurrentTime string
	IsComplate  bool `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
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

func FromDomain(domain *materies.Domain) *Materi {
	return &Materi{
		ID:        domain.ID,
		Title:     domain.Title,
		ModulID:   domain.ModulId,
		Order:     int8(domain.Order),
		Type:      TYPE(domain.Type),
		UrlMateri: domain.UrlMateri,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (r Materi) ToDomain() materies.Domain {
	var typeStr string
	if r.Type == 1 {
		typeStr = "text"
	} else {
		typeStr = "video"
	}
	return materies.Domain{
		ID:         r.ID,
		Title:      r.Title,
		ModulId:    r.ModulID,
		Order:      int(r.Order),
		Type:       int(r.Type),
		TypeString: typeStr,
		UrlMateri:  r.UrlMateri,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}
