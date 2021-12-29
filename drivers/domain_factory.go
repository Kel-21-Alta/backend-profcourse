package drivers

import (
	"gorm.io/gorm"
	_coursesUsecase "profcourse/business/courses"
	_localsRepository "profcourse/business/locals"
	_modulsUsecase "profcourse/business/moduls"
	"profcourse/business/smtp_email"
	_userUsecase "profcourse/business/users"
	_usersCoursesUsecase "profcourse/business/users_courses"
	"profcourse/drivers/databases/courses"
	"profcourse/drivers/databases/moduls"
	"profcourse/drivers/databases/users"
	"profcourse/drivers/databases/users_courses"
	"profcourse/drivers/locals"
	"profcourse/drivers/thirdparties/smtp"
)

func NewMysqlUserRepository(conn *gorm.DB) _userUsecase.Repository {
	return users.NewMysqlRepository(conn)
}

func NewSmtpRepository(config smtp.SmtpEmail) smtp_email.Repository {
	return smtp.NewSmtpEmail(config)
}

func NewLocalRepository() _localsRepository.Repository {
	return locals.NewLocalRepository()
}

func NewMysqlCourseRepository(conn *gorm.DB) _coursesUsecase.Repository {
	return courses.NewMysqlRepository(conn)
}

func NewMysqlUserCourseRepository(conn *gorm.DB) _usersCoursesUsecase.Repository {
	return users_courses.NewMysqlRepository(conn)
}

func NewMysqlModulRepository(conn *gorm.DB) _modulsUsecase.Repository {
	return moduls.NewMysqlRepository(conn)
}
