package users_spesializations

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"profcourse/business/users_spesializations"
	"time"
)

type UsersSpesializations struct {
	ID               string `gorm:"primaryKey;not null;unique"`
	UserID           string `gorm:"not null"`
	SpesializationID string `gorm:"not null"`
	Progress         int    `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

func (c *UsersSpesializations) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now().Local()
	return nil
}

func (c *UsersSpesializations) BeforeUpdate(db *gorm.DB) error {
	c.UpdatedAt = time.Now().Local()
	return nil
}


func FromDomain(domain *users_spesializations.Domain) UsersSpesializations {
	return UsersSpesializations{
		ID:               domain.ID,
		UserID:           domain.UserID,
		SpesializationID: domain.SpesializationID,
		Progress:         domain.Progress,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func (s UsersSpesializations) ToDomain() users_spesializations.Domain {
	return users_spesializations.Domain{
		ID:               s.ID,
		UserID:           s.UserID,
		SpesializationID: s.SpesializationID,
		Progress:         s.Progress,
		CreatedAt:        s.CreatedAt,
		UpdatedAt:        s.UpdatedAt,
	}
}
