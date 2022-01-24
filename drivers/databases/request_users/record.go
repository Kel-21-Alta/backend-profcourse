package requestusers

import (
	"profcourse/business/request_users"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type RequestUser struct {
	ID                string `gorm:"primaryKey;unique;not null"`
	UserId            string `gorm:"not null"`
	CategoryRequestId string `gorm:"size:191"`
	Request           string `gorm:"not null"`

	CategoryRequest CategoryRequest `gorm:"foreignKey:CategoryRequestId"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type CategoryRequest struct {
	ID        string `gorm:"primaryKey;unique;not null"`
	Title     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *RequestUser) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4().String()
	u.CreatedAt = time.Now().Local()

	return
}

func (u *RequestUser) BeforeUpdate(db *gorm.DB) error {
	u.UpdatedAt = time.Now().Local()
	return nil
}

func (u *CategoryRequest) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4().String()
	u.CreatedAt = time.Now().Local()

	return
}

func (u *CategoryRequest) BeforeUpdate(db *gorm.DB) error {
	u.UpdatedAt = time.Now().Local()
	return nil
}

func FromDomain(domain *request_users.Domain) *RequestUser {
	return &RequestUser{
		ID:                domain.Id,
		UserId:            domain.UserId,
		CategoryRequestId: domain.CategoryID,
		Request:           domain.Request,
		CategoryRequest: CategoryRequest{
			ID:        domain.Category.ID,
			Title:     domain.Category.Title,
			CreatedAt: domain.Category.CreatedAt,
			UpdatedAt: domain.Category.UpdatedAt,
		},
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (r *RequestUser) ToDomain() request_users.Domain {
	return request_users.Domain{
		Id:         r.ID,
		UserId:     r.UserId,
		CategoryID: r.CategoryRequestId,
		Request:    r.Request,
		Category: request_users.Category{
			ID:        r.CategoryRequest.ID,
			Title:     r.CategoryRequest.Title,
			CreatedAt: r.CategoryRequest.CreatedAt,
			UpdatedAt: r.CategoryRequest.UpdatedAt,
		},
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func ToListCategoryDomain(recs []CategoryRequest) []request_users.Category {
	var listCategory []request_users.Category

	for _, rec := range recs {
		listCategory = append(listCategory, request_users.Category{
			ID:        rec.ID,
			Title:     rec.Title,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
		})
	}

	return listCategory
}
