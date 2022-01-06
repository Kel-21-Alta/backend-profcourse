package users_courses

import (
	"context"
	"gorm.io/gorm"
	_usersCoursesUsecase "profcourse/business/users_courses"
)

type msqlUserCourseRepository struct {
	Conn *gorm.DB
}

func (m msqlUserCourseRepository) GetEndRollCourseUserById(ctx context.Context, domain *_usersCoursesUsecase.Domain) (*_usersCoursesUsecase.Domain, error) {
	var err error
	rec := UsersCourses{}

	err = m.Conn.Where("course_id = ? AND user_id = ?", domain.CourseId, domain.UserId).First(&rec).Error
	if err != nil {
		return &_usersCoursesUsecase.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (m msqlUserCourseRepository) UserRegisterCourse(ctx context.Context, domain *_usersCoursesUsecase.Domain) (*_usersCoursesUsecase.Domain, error) {
	err := m.Conn.Create(FromDomain(*domain)).Error
	if err != nil {
		return &_usersCoursesUsecase.Domain{}, err
	}

	return domain, nil
}
func NewMysqlRepository(conn *gorm.DB) *msqlUserCourseRepository {
	return &msqlUserCourseRepository{Conn: conn}
}
