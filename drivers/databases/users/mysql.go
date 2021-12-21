package users

import (
	"context"
	"gorm.io/gorm"
	"profcourse/business/users"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func (m mysqlUserRepository) CreateUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	rec := FromDomain(domain)
	result := m.Conn.Create(&rec)
	if result.Error != nil {
		return rec.ToDomain(), result.Error
	}
	return rec.ToDomain(), nil
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &mysqlUserRepository{
		Conn: conn,
	}
}
