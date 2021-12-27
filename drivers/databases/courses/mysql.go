package courses

import (
	"context"
	"gorm.io/gorm"
	"profcourse/business/courses"
)

type mysqlCourseRepository struct {
	Conn *gorm.DB
}

func (r mysqlCourseRepository) GetAllCourses(ctx context.Context, domain *courses.Domain) (*[]courses.Domain, error) {
	var coursesResult []Courses

	err := r.Conn.Find(&coursesResult).Error
	if err != nil {
		return &[]courses.Domain{}, err
	}

	return ToListDomain(coursesResult), nil
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
