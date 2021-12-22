package drivers

import (
	"gorm.io/gorm"
	_userUsecase "profcourse/business/users"
	"profcourse/drivers/databases/users"
)

func NewMysqlUserRepository(conn *gorm.DB) _userUsecase.Repository {
	return users.NewMysqlRepository(conn)
}
