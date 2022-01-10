package courses

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"profcourse/business/courses"
	controller "profcourse/controllers"
)

type mysqlCourseRepository struct {
	Conn *gorm.DB
}

func (r *mysqlCourseRepository) DeleteCourseForUser(ctx context.Context, id string, token courses.Token) (courses.Domain, error) {
	var rec Courses
	var err error
	err = r.Conn.First(&rec, "id = ?", id).Error
	if rec.ID != token.UserId {
		return courses.Domain{}, controller.FORBIDDIN_USER
	}
	err = r.Conn.Delete(&rec).Error
	if err != nil {
		return courses.Domain{}, err
	}
	return courses.Domain{}, nil
}

func (r *mysqlCourseRepository) DeleteCourseForAdmin(ctx context.Context, id string) (courses.Domain, error) {
	var rec Courses
	err := r.Conn.Delete(&rec, "id = ?", id).Error
	if err != nil {
		return courses.Domain{}, err
	}
	return courses.Domain{}, nil
}

func (r *mysqlCourseRepository) UpdateCourseForUser(ctx context.Context, domain *courses.Domain, token *courses.Token) (courses.Domain, error) {
	var rec Courses
	var err error
	err = r.Conn.First(&rec, "id = ?", domain.ID).Error

	if err != nil {
		return courses.Domain{}, err
	}

	if rec.TeacherId != token.UserId {
		return courses.Domain{}, controller.FORBIDDIN_USER
	}

	rec.Title = domain.Title
	rec.Description = domain.Description
	rec.ImgUrl = domain.Description

	err = r.Conn.Save(&rec).Error
	if err != nil {
		return courses.Domain{}, err
	}

	return *rec.ToDomain(), err
}

func (r *mysqlCourseRepository) UpdateCourseForAdmin(ctx context.Context, domain *courses.Domain) (courses.Domain, error) {
	var rec Courses
	var err error
	err = r.Conn.First(&rec, "id = ?", domain.ID).Error

	if err != nil {
		return courses.Domain{}, err
	}

	rec.Title = domain.Title
	rec.Description = domain.Description
	rec.ImgUrl = domain.Description

	err = r.Conn.Save(&rec).Error
	if err != nil {
		return courses.Domain{}, err
	}

	return *rec.ToDomain(), err
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
	count := 0

	err := r.Conn.Preload(clause.Associations).First(&rec, " id = ?", domain.ID).Error
	if err != nil {
		return &courses.Domain{}, err
	}

	resultDomain := rec.ToDomain()
	var RegisteredUsers []RegisteredUser

	// Mendapatkan jumlah user yang mengambil course
	r.Conn.Table("courses").Select("COUNT(users_courses.id) as count").Joins("INNER JOIN users_courses ON users_courses.course_id = courses.id").Where("courses.id = ?", resultDomain.ID).Scan(&count)

	resultDomain.UserTakenCourse = count

	// Mendapatkan list user yang mengambil course untuk rangking
	r.Conn.Table("courses").Select(" users_courses.user_id  as user_id, users.name as name_user, users_courses.skor as skor, users_courses.progress as progress").Joins("INNER JOIN users_courses ON users_courses.course_id = courses.id").Joins("INNER JOIN users ON users_courses.user_id = users.id").Where("courses.id = ?", resultDomain.ID).Order("users_courses.skor desc").Limit(10).Scan(&RegisteredUsers)

	resultDomain.InfoUser = domain.InfoUser

	result := FromRegiteredUserToDomain(resultDomain, RegisteredUsers)
	return result, err
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
