package getOneRequestUser

import (
	"profcourse/business/request_users"
	"time"
)

type GetOneRequestUserResponse struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	CategoryID string `json:"category_id"`
	Topik string `json:"topik"`
	CategoryName string `json:"category_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain request_users.Domain) *GetOneRequestUserResponse {
	return &GetOneRequestUserResponse{
		Id:           domain.Id,
		UserId:       domain.UserId,
		CategoryID:   domain.CategoryID,
		Topik:        domain.Request,
		CategoryName: domain.Category.Title,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
