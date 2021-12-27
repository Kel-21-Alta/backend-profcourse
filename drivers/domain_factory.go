package drivers

import (
	"gorm.io/gorm"
	_coursesUsecase "profcourse/business/courses"
	_localsRepository "profcourse/business/locals"
	"profcourse/business/smtp_email"
	_userUsecase "profcourse/business/users"
	"profcourse/drivers/databases/courses"
	"profcourse/drivers/databases/users"
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
