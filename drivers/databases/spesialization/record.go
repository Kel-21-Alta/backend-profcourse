package spesialization

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"profcourse/business/spesializations"
	"profcourse/drivers/databases/courses"
	"time"
)

type Spesialization struct {
	ID            string `gorm:"primaryKey;unique"`
	Title         string `gorm:"not null;"`
	Description   string `gorm:"not null"`
	ImageUrl      string `gorm:"not null"`
	CertificateId string

	Courses []*courses.Courses `gorm:"many2many:spesialization_courses;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (c *Spesialization) BeforeCreate(db *gorm.DB) error {
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now().Local()
	return nil
}

func (c *Spesialization) BeforeUpdate(db *gorm.DB) error {
	c.UpdatedAt = time.Now().Local()
	return nil
}

func FromDomain(domain *spesializations.Domain) *Spesialization {

	var listCourses []*courses.Courses

	for _, course := range domain.CourseIds {
		listCourses = append(listCourses, &courses.Courses{ID: course})
	}

	return &Spesialization{
		ID:            domain.ID,
		Title:         domain.Title,
		Description:   domain.Description,
		ImageUrl:      domain.ImageUrl,
		CertificateId: domain.CertificateId,
		Courses:       listCourses,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func (s *Spesialization) ToDomain() spesializations.Domain {
	var listCourseId []string
	return spesializations.Domain{
		ID:            s.ID,
		Title:         s.Title,
		ImageUrl:      s.ImageUrl,
		CourseIds:     listCourseId,
		Description:   s.Description,
		CreatedAt:     s.CreatedAt,
		UpdatedAt:     s.UpdatedAt,
		CertificateId: s.CertificateId,
	}
}
func (s *Spesialization) ToDomainWithCourses() spesializations.Domain {
	var list []spesializations.Course
	for _, courseRec := range s.Courses {
		list = append(list, spesializations.Course{
			ID:          courseRec.ID,
			Title:       courseRec.Title,
			Rating:      0,
			Description: courseRec.Description,
		})
	}
	return spesializations.Domain{
		ID:            s.ID,
		Title:         s.Title,
		ImageUrl:      s.ImageUrl,
		Description:   s.Description,
		CreatedAt:     s.CreatedAt,
		UpdatedAt:     s.UpdatedAt,
		CertificateId: s.CertificateId,
		Courses:       list,
	}
}

func ToListDomain(s []*Spesialization) []spesializations.Domain {
	var list []spesializations.Domain

	for _, spesialization := range s {
		list = append(list, spesialization.ToDomain())
	}
	return list
}
