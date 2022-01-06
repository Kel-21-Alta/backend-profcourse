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

func (r *mysqlCourseRepository) GetCountCourse(ctx context.Context) (*courses.Summary, error) {
	result := 0
	err := r.Conn.Raw("SELECT COUNT(*) as result FROM courses").Scan(&result).Error
	if err != nil {
		return &courses.Summary{}, err
	}
	return &courses.Summary{CountCourse: result}, nil
}

func (r *mysqlCourseRepository) GetOneCourse(ctx context.Context, domain *courses.Domain) (*courses.Domain, error) {
	rec := Courses{}

	err := r.Conn.Preload(clause.Associations).First(&rec, " id = ?", domain.ID).Error
	if err != nil {
		return &courses.Domain{}, err
	}

	return rec.ToDomain(), err
}

// Paginate Fungsi ini untuk mengimplementasikan pagination pada list course
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

// GetAllCourses Untuk mendapatkan semua list course sesuai keperluan
func (r *mysqlCourseRepository) GetAllCourses(ctx context.Context, domain *courses.Domain) (*[]courses.Domain, error) {
	var coursesResult []*Courses
	var err error
	err = r.Conn.Scopes(Paginate(*domain)).Order(domain.Sort+" "+domain.SortBy).Where("title Like ?", "%"+domain.KeywordSearch+"%").Find(&coursesResult).Error

	if err != nil {
		return &[]courses.Domain{}, err
	}

	return ToListDomain(coursesResult), nil
}

func (r *mysqlCourseRepository) CreateCourse(ctx context.Context, domain *courses.Domain) (*courses.Domain, error) {
	rec := FromDomain(*domain)

	err := r.Conn.Create(&rec).Error
	if err != nil {
		return &courses.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) courses.Repository {
	return &mysqlCourseRepository{
		Conn: conn,
	}
}
