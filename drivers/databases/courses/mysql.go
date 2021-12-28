package courses

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"profcourse/business/courses"
)

type mysqlCourseRepository struct {
	Conn *gorm.DB
}

func (r mysqlCourseRepository) GetOneCourse(ctx context.Context, domain *courses.Domain) (*courses.Domain, error) {
	rec := Courses{}

	err := r.Conn.Preload(clause.Associations).First(&rec, domain.ID).Error
	if err != nil {
		return &courses.Domain{}, err
	}

	return rec.ToDomain(), err
}

func (r mysqlCourseRepository) CreateCourse(ctx context.Context, domain *courses.Domain) (*courses.Domain, error) {
	rec := FromDomain(*domain)

	err := r.Conn.Create(&rec).Error
	if err != nil {
		return &courses.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) *mysqlCourseRepository {
	return &mysqlCourseRepository{
		Conn: conn,
	}
}
