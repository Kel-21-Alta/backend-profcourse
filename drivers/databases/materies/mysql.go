package materies

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/materies"
)

type MateriesRepository struct {
	Conn *gorm.DB
}

func (m MateriesRepository) UpdateMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {

	var rec Materi
	var err error

	err = m.Conn.First(&rec, "id = ?", domain.ID).Error

	if err != nil {
		return materies.Domain{}, err
	}

	rec.Title = domain.Title
	rec.UrlMateri = domain.UrlMateri
	rec.ModulID = domain.ModulId
	rec.Type = TYPE(domain.Type)
	rec.Order = int8(domain.Order)

	err = m.Conn.Save(&rec).Error

	if err != nil {
		return materies.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (m MateriesRepository) DeleteMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {
	var rec Materi
	err := m.Conn.Delete(&rec, "id = ?", domain.ID).Error

	if err != nil {
		return materies.Domain{}, err
	}

	return *domain, nil
}

func (m MateriesRepository) CreateMateri(ctx context.Context, domain *materies.Domain) (materies.Domain, error) {
	var req = FromDomain(domain)

	err := m.Conn.Create(&req).Error
	if err != nil {
		return materies.Domain{}, err
	}

	return req.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) materies.Repository {
	return &MateriesRepository{Conn: conn}
}
