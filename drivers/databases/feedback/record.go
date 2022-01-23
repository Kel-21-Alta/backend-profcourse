package feedback

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"profcourse/business/feedback"
	"profcourse/drivers/databases/courses"
	"profcourse/drivers/databases/users"
	"time"
)

type Feedback struct {
	ID       string `gorm:"primaryKey;unique;not null"`
	UserId   string `gorm:"not null;size:191;index:idx_unique3,unique"`
	CourseId string `gorm:"not null;size:191;index:idx_unique3,unique"`
	Review   string
	Rating   float32

	User   users.User      `gorm:"foreignKey:UserId"`
	Course courses.Courses `gorm:"foreignKey:CourseId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *Feedback) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewV4().String()
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *Feedback) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}

func FromDomain(domain *feedback.Domain) *Feedback {
	return &Feedback{
		ID:        domain.ID,
		UserId:    domain.UserId,
		CourseId:  domain.CourseId,
		Review:    domain.Review,
		Rating:    domain.Rating,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdateAt,
	}
}

func (f *Feedback) ToDomain() feedback.Domain {
	return feedback.Domain{
		ID:        f.ID,
		UserId:    f.UserId,
		CourseId:  f.CourseId,
		Review:    f.Review,
		Rating:    f.Rating,
		CreatedAt: f.CreatedAt,
		UpdateAt:  f.UpdatedAt,
	}
}
