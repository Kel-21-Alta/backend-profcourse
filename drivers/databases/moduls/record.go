package moduls

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Moduls struct {
	ID       string `gorm:"primaryKey;unique"`
	Title    string `gorm:"not null"`
	Order    int    `gorm:"not null"`
	CourseId string `gorm:"not null;size:191"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}


func (c *Moduls) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now().Local()
	return nil
}

func (c *Moduls) BeforeUpdate(db *gorm.DB) error {
	c.UpdatedAt = time.Now().Local()
	return nil
}

// TODO: Sampai sini buat TODOMAIN dan FROM DOMAIN