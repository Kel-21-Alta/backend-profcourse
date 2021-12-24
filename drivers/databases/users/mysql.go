package users

import (
	"context"
	"gorm.io/gorm"
	"profcourse/business/users"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func (m mysqlUserRepository) GetUserById(ctx context.Context, id string) (users.Domain, error) {
	rec := User{}
	err := m.Conn.First(&rec, "id = ?", id).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

/*
	Digunakan untuk mendapatkan user dengan email tertentu
**/
func (m mysqlUserRepository) GetUserByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := User{}
	err := m.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

// Digunakan untuk membuat user baru
func (m mysqlUserRepository) CreateUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	rec := FromDomain(domain)
	result := m.Conn.Create(&rec)
	if result.Error != nil {
		return rec.ToDomain(), result.Error
	}
	return rec.ToDomain(), nil
}

// untuk mengupdate password
func (m mysqlUserRepository) UpdatePassword(ctx context.Context, domain users.Domain, hash string) (users.Domain, error) {
	rec := FromDomain(domain)
	err := m.Conn.Model(&rec).Update("password", hash).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) users.Repository {
	return &mysqlUserRepository{
		Conn: conn,
	}
}
