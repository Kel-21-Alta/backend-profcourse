package spesialization

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"profcourse/business/spesializations"
)

type spesializationRepository struct {
	Conn *gorm.DB
}

func (s spesializationRepository) CreateSpasialization(ctx context.Context, domain *spesializations.Domain) (*spesializations.Domain, error) {

	req := FromDomain(domain)

	err := s.Conn.Omit("Courses.*").Create(&req).Error

	if err != nil {
		return &spesializations.Domain{}, err
	}

	return req.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) spesializations.Repository {
	return &spesializationRepository{Conn: conn}
}
