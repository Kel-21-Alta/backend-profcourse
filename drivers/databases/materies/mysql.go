package materies

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/materies"
)

type MateriesRepository struct {
	Conn *gorm.DB
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
