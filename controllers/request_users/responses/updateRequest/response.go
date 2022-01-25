package updateRequest

import (
	"profcourse/business/request_users"
	"time"
)

type UpdateRequestResponse struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	CategoryID string `json:"category_id"`
	Topik string `json:"topik"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain request_users.Domain) *UpdateRequestResponse {
	return &UpdateRequestResponse{
		Id:         domain.Id,
		UserId:     domain.UserId,
		CategoryID: domain.CategoryID,
		Topik:      domain.Request,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
