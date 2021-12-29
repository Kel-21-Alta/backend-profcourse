package moduls

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/moduls"
)

type mysqlModulsRepository struct {
	Conn *gorm.DB
}

func (m mysqlModulsRepository) CreateModul(ctx context.Context, domain *moduls.Domain) (*moduls.Domain, error) {
	var err error
	rec := FromDomain(domain)

	err = m.Conn.Create(&rec).Error
	if err != nil {
		return &moduls.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) *mysqlModulsRepository {
	return &mysqlModulsRepository{Conn: conn}
}
