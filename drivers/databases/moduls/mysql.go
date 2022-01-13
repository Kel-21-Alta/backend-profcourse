package moduls

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/moduls"
)

type mysqlModulsRepository struct {
	Conn *gorm.DB
}

func (m mysqlModulsRepository) UpdateModul(ctx context.Context, domain *moduls.Domain) (moduls.Domain, error) {
	var rec Moduls
	var err error

	err = m.Conn.First(&rec, "id = ?", domain.ID).Error
	if err != nil {
		return moduls.Domain{}, err
	}

	rec.Title = domain.Title
	rec.Order = domain.Order

	err = m.Conn.Save(&rec).Error
	if err != nil {
		return moduls.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (m mysqlModulsRepository) GetOneModul(ctx context.Context, domain *moduls.Domain) (moduls.Domain, error) {
	var rec Moduls
	err := m.Conn.Preload("Materies").First(&rec, "id = ?", domain.ID).Error

	if err != nil {
		return moduls.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (m mysqlModulsRepository) CreateModul(ctx context.Context, domain *moduls.Domain) (moduls.Domain, error) {
	var err error
	rec := FromDomain(domain)

	err = m.Conn.Create(&rec).Error
	if err != nil {
		return moduls.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) moduls.Repository {
	return mysqlModulsRepository{Conn: conn}
}
