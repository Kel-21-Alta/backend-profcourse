package registerSpesialization

import (
	"profcourse/business/users_spesializations"
	"time"
)

type RegisterSpesializationResponse struct {
	ID string `json:"id"`
	UserId string `json:"user_id"`
	SpesializationId string `json:"spesialization_id"`
	Progress int `json:"progress"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users_spesializations.Domain) *RegisterSpesializationResponse {
	return &RegisterSpesializationResponse{
		ID:               domain.ID,
		UserId:           domain.UserID,
		SpesializationId: domain.SpesializationID,
		Progress:         domain.Progress,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}