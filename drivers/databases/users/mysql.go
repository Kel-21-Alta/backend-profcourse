package users

import (
	"context"
	"gorm.io/gorm"
	"profcourse/business/users"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

/*
	Digunakan untuk mendapatkan User dengan email tertentu
**/
func (m mysqlUserRepository) GetUserByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := User{}
	err := m.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (m mysqlUserRepository) CreateUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	rec := FromDomain(domain)
	result := m.Conn.Create(&rec)
	if result.Error != nil {
		return rec.ToDomain(), result.Error
	}
	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) users.Repository {
	return &mysqlUserRepository{
		Conn: conn,
	}
}
