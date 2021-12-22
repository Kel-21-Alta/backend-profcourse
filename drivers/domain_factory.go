package drivers

import (
	"gorm.io/gorm"
	"profcourse/business/smtpEmail"
	_userUsecase "profcourse/business/users"
	"profcourse/drivers/databases/users"
	"profcourse/drivers/thirdparties/smtp"
)

func NewMysqlUserRepository(conn *gorm.DB) _userUsecase.Repository {
	return users.NewMysqlRepository(conn)
}

func NewSmtpRepository(config smtp.SmtpEmail) smtpEmail.Repository {
	return smtp.NewSmtpEmail(config)
}
