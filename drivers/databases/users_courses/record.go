package users_courses

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"profcourse/business/users_courses"
	"time"
)

type UsersCourses struct {
	ID          string `gorm:"primaryKey;unique"`
	UserId      string `gorm:"size:191;uniqueIndex:idx_user_course"`
	CourseId    string `gorm:"size:191;uniqueIndex:idx_user_course"`
	Progress    int
	LastVideoId string `gorm:"size:191"`
	LastModulId string `gorm:"size:191"`
	Skor        int    `gorm:"default:0"`

	//User   users.User      `gorm:"foreignKey:UserMakeModul;references:ID"`
	//Course courses.Courses `gorm:"foreignKey:CourseId;references:ID"`

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
		Score:       u.Skor,
	}
}

func FromDomain(domain users_courses.Domain) *UsersCourses {
	return &UsersCourses{
		UserId:      domain.UserId,
		CourseId:    domain.CourseId,
		Progress:    0,
		LastVideoId: domain.LastVideoId,
		LastModulId: domain.LastModulId,
		Skor:        domain.Score,
	}
}

func ToUserDomain(user_course []UsersCourses) users_courses.User {

	var listCourse []users_courses.Domain

	for _, course := range user_course {
		listCourse = append(listCourse, *course.ToDomain())
	}

	return users_courses.User{
		UserID:      "",
		Name:        "",
		CountCourse: len(user_course),
		Courses:     listCourse,
	}

}
