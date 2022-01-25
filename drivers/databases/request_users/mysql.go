package requestusers

import (
	"context"
	"profcourse/business/request_users"

	"gorm.io/gorm"
)

type RequestUserRepo struct {
	Conn *gorm.DB
}

func (r *RequestUserRepo) GetOneRequestUser(ctx context.Context, domain *request_users.Domain) (request_users.Domain, error) {
	var rec RequestUser
	err := r.Conn.Preload("CategoryRequest").First(&rec, "id = ?", domain.Id).Error

	if err != nil {
		return request_users.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (r *RequestUserRepo) UpdateRequestUser(ctx context.Context, domain *request_users.Domain) (request_users.Domain, error) {
	var rec RequestUser
	var err error

	err = r.Conn.Where("user_id = ?", domain.UserId).First(&rec, "id = ?", domain.Id).Error
	if err != nil {
		return request_users.Domain{}, err
	}

	rec.Request = domain.Request
	rec.CategoryRequestId = domain.CategoryID

	err = r.Conn.Save(&rec).Error

	if err != nil {
		return request_users.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (r *RequestUserRepo) DeleteRequestUser(ctx context.Context, domain *request_users.Domain) (request_users.Domain, error) {
	var rec RequestUser
	err := r.Conn.Delete(&rec, "id = ?", domain.Id).Error
	if err != nil {
		return request_users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

// Paginate Fungsi ini untuk mengimplementasikan pagination pada list course
func Paginate(domain request_users.Domain) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := domain.Query.Offset
		limit := domain.Query.Limit
		if limit == 0 {
			limit = 10
		}
		return db.Offset(offset).Limit(limit)
	}
}

func (r *RequestUserRepo) GetAllRequestUser(ctx context.Context, domain *request_users.Domain) ([]request_users.Domain, error) {
	var recs []RequestUser
	var err error
	err = r.Conn.Preload("CategoryRequest").Scopes(Paginate(*domain)).Order("created_at "+domain.Query.Sort).Where("user_id = ?", domain.UserId).Where("request Like ?", "%"+domain.Query.Search+"%").Find(&recs).Error

	if err != nil {
		return []request_users.Domain{}, err
	}

	return ToListRequestUserDomain(recs), nil
}

func (r *RequestUserRepo) AdminGetAllRequestUser(ctx context.Context, domain *request_users.Domain) ([]request_users.Domain, error) {
	var recs []RequestUser
	var err error
	err = r.Conn.Preload("CategoryRequest").Preload("User").Scopes(Paginate(*domain)).Order("created_at "+domain.Query.Sort).Where("request Like ?", "%"+domain.Query.Search+"%").Find(&recs).Error

	if err != nil {
		return []request_users.Domain{}, err
	}

	return ToListRequestUserDomain(recs), nil
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
