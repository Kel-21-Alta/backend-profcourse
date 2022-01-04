package drivers

import (
	"gorm.io/gorm"
	_coursesUsecase "profcourse/business/courses"
	_modulsUsecase "profcourse/business/moduls"
	"profcourse/business/send_email"
	_localsRepository "profcourse/business/uploads"
	_userUsecase "profcourse/business/users"
	_usersCoursesUsecase "profcourse/business/users_courses"
	"profcourse/drivers/databases/courses"
	"profcourse/drivers/databases/moduls"
	"profcourse/drivers/databases/users"
	"profcourse/drivers/databases/users_courses"
	"profcourse/drivers/thirdparties/mail"
	"profcourse/drivers/uploadLocals"
)

func NewMysqlUserRepository(conn *gorm.DB) _userUsecase.Repository {
	return users.NewMysqlRepository(conn)
}

func NewSmtpRepository(config mail.Email) send_email.Repository {
	return mail.NewSmtpEmail(config)
}

func NewLocalRepository() _localsRepository.Repository {
	return uploadLocals.NewLocalRepository()
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
