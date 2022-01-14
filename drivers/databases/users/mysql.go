package users

import (
	"context"
	"profcourse/business/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func (m mysqlUserRepository) UpdateDataCurrentUser(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	var rec User
	var err error

	err = m.Conn.First(&rec, "id = ?", domain.ID).Error
	if err != nil {
		return users.Domain{}, err
	}

	rec.ImgProfile = domain.ImgProfile
	rec.Birth = domain.Birth
	rec.NoHp = domain.NoHp
	rec.Name = domain.Name
	rec.Bio = domain.Bio
	rec.BirthPlace = domain.BirthPlace

	err = m.Conn.Save(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (m mysqlUserRepository) GetCountUser(ctx context.Context) (*users.Summary, error) {
	var result int
	err := m.Conn.Raw("SELECT COUNT(*) as result FROM users").Scan(&result).Error
	if err != nil {
		return &users.Summary{}, nil
	}
	return &users.Summary{CountUser: result}, nil
}

func (m mysqlUserRepository) DeleteUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	rec := User{}
	result := m.Conn.Where("id = ?", domain.IdUser).Delete(&rec)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	domain.Message = "User dengan id: " + domain.IdUser + " telah dihapus"
	return domain, nil
}

// UpdateUser Update data user dari admin
func (m mysqlUserRepository) UpdateUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	first := User{}
	err := m.Conn.First(&first, "id = ?", domain.ID).Error
	first.Name = domain.Name
	first.NoHp = domain.NoHp
	first.Bio = domain.Bio
	first.Birth = domain.Birth
	first.BirthPlace = domain.BirthPlace
	result := m.Conn.Save(&first)
	if err != nil {
		return users.Domain{}, err
	}
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return first.ToDomain(), nil
}

// GetUserById Untuk mendapatkan data user berdasarkan id
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
