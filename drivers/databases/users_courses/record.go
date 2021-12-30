package users_courses

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"profcourse/business/users_courses"
	"profcourse/drivers/databases/courses"
	"profcourse/drivers/databases/users"
	"time"
)

type UsersCourses struct {
	ID          string `gorm:"primaryKey;unique"`
	UserId      string `gorm:"size:191;uniqueIndex:idx_user_course"`
	CourseId    string `gorm:"size:191;uniqueIndex:idx_user_course"`
	Progress    int
	LastVideoId string `gorm:"size:191"`
	LastModulId string `gorm:"size:191"`

	User   users.User      `gorm:"foreignKey:UserId;references:ID"`
	Course courses.Courses `gorm:"foreignKey:CourseId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *UsersCourses) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4().String()
	u.CreatedAt = time.Now().Local()

	return
}

func (u *UsersCourses) BeforeUpdate(db *gorm.DB) error {
	u.UpdatedAt = time.Now().Local()
	return nil
}

func (u UsersCourses) ToDomain() *users_courses.Domain {
	return &users_courses.Domain{
		ID:          u.ID,
		UserId:      u.UserId,
		CourseId:    u.CourseId,
		Progres:     u.Progress,
		LastVideoId: u.LastVideoId,
		LastModulId: u.LastModulId,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.CreatedAt,
	}
}

func FromDomain(domain users_courses.Domain) *UsersCourses {
	return &UsersCourses{
		UserId:      domain.UserId,
		CourseId:    domain.CourseId,
		Progress:    0,
		LastVideoId: domain.LastVideoId,
		LastModulId: domain.LastModulId,
	}
}