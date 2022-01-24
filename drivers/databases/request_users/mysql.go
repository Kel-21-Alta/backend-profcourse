package requestusers

import (
	"context"
	"profcourse/business/request_users"

	"gorm.io/gorm"
)

type RequestUserRepo struct {
	Conn *gorm.DB
}

func (r *RequestUserRepo) GetAllCategoryRequest(ctx context.Context) ([]request_users.Category, error) {
	var rec []CategoryRequest

	err := r.Conn.Find(&rec).Error
	if err != nil {
		return []request_users.Category{}, err
	}

	return ToListCategoryDomain(rec), nil
}

func (r *RequestUserRepo) GetOneRequest(ctx context.Context, domain *request_users.Domain) (request_users.Domain, error) {
	var rec = FromDomain(domain)

	err := r.Conn.Preload("CategoryRequest").First(&rec, "id = ?", rec.ID).Error

	if err != nil {
		return request_users.Domain{}, err
	}
	return rec.ToDomain(), nil
}


func (r *RequestUserRepo) CreateRequest(ctx context.Context, domain *request_users.Domain) (request_users.Domain, error) {

	var rec = FromDomain(domain)

	err := r.Conn.Create(&rec).Error

	if err != nil {
		return request_users.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func NewMysqlRepository(conn *gorm.DB) request_users.Repository {
	return &RequestUserRepo{
		Conn: conn,
	}
}
