package courses

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"profcourse/business/courses"
	"profcourse/drivers/databases/moduls"
	"profcourse/drivers/databases/users"
	"time"
)

type STATUS int8

const (
	Publish STATUS = 1
	Draft   STATUS = 2
	Pending STATUS = 3
)

type Courses struct {
	ID          string `gorm:"primaryKey;unique"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	ImgUrl      string
	TeacherId   string `gorm:"size:191"`
	Status      STATUS `gorm:"default:2"`
	StatusText  string

	Teacher users.User `gorm:"foreignKey:TeacherId;references:ID"`

	Moduls []moduls.Moduls `gorm:"foreignKey:CourseId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (c *Courses) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now().Local()
	return nil
}

func (c *Courses) BeforeUpdate(db *gorm.DB) error {
	c.UpdatedAt = time.Now().Local()
	return nil
}

func FromDomain(domain courses.Domain) *Courses {
	return &Courses{
		ID:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		ImgUrl:      domain.ImgUrl,
		TeacherId:   domain.TeacherId,
		Status:      STATUS(domain.Status),
		StatusText:  domain.StatusText,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func (c Courses) ToDomain() *courses.Domain {

	var listModuls []courses.Modul
	for _, modul := range c.Moduls {
		listModuls = append(listModuls, courses.Modul{
			NameModul: modul.Title,
			ModulID:   modul.ID,
			Order: modul.Order,
		})
	}

	return &courses.Domain{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		ImgUrl:      c.ImgUrl,
		TeacherId:   c.TeacherId,
		Status:      int8(c.Status),
		StatusText:  c.StatusText,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		TeacherName: c.Teacher.Name,
		Moduls:      listModuls,
	}
}

func ToListDomain(listCourses []*Courses) *[]courses.Domain {
	var listCourseDomain []courses.Domain
	for _, course := range listCourses {
		listCourseDomain = append(listCourseDomain, *course.ToDomain())
	}
	return &listCourseDomain
}
