package users_spesializations

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/users_spesializations"
)

type UsersSpesializationsRepository struct {
	Conn *gorm.DB
}

func (u UsersSpesializationsRepository) GetEndRollSpesializationById(ctx context.Context, domain *users_spesializations.Domain) (users_spesializations.Domain, error) {
	var err error
	var rec UsersSpesializations

	err = u.Conn.Where("user_id = ? AND spesialization_id = ?", domain.UserID, domain.SpesializationID).First(&rec).Error

	if err != nil {
		return users_spesializations.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (u UsersSpesializationsRepository) RegisterSpesialization(ctx context.Context, domain *users_spesializations.Domain) (users_spesializations.Domain, error) {

	var rec UsersSpesializations = FromDomain(domain)

	err := u.Conn.Create(&rec).Error

	if err != nil {
		return users_spesializations.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func NewMyslRepository(conn *gorm.DB) users_spesializations.Repository {
	return &UsersSpesializationsRepository{Conn: conn}
}
