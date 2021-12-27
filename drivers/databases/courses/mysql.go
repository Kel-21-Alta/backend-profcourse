package courses

import (
	"context"
	"gorm.io/gorm"
	"profcourse/business/courses"
)

type mysqlCourseRepository struct {
	Conn *gorm.DB
}

// Fungsi ini untuk mengimplementasikan pagination pada list course
func Paginate(domain courses.Domain) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := domain.Offset
		limit := domain.Limit
		if limit == 0 {
			limit = 10
		}
		return db.Offset(offset).Limit(limit)
	}
}

// Untuk mendapatkan semua list course sesuai keperluan
func (r mysqlCourseRepository) GetAllCourses(ctx context.Context, domain *courses.Domain) (*[]courses.Domain, error) {
	var coursesResult []Courses
	var err error
	err = r.Conn.Scopes(Paginate(*domain)).Order(domain.Sort+" "+domain.SortBy).Where("title Like ?", "%"+domain.KeywordSearch+"%").Find(&coursesResult).Error

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
