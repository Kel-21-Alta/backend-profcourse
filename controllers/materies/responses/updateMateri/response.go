package updateMateri

import (
	"profcourse/business/materies"
	"time"
)

type UpdateMateriResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	UrlMateri string    `json:"url_materi"`
	Order     int       `json:"order"`
	ModulId   string    `json:"modul_id"`
	CreatedId time.Time `json:"created_id"`
	UpdatedId time.Time `json:"updated_id"`
}

func FromDomain(domain materies.Domain) *UpdateMateriResponse {
	return &UpdateMateriResponse{
		ID:        domain.ID,
		Title:     domain.Title,
		Type:      domain.TypeString,
		UrlMateri: domain.UrlMateri,
		Order:     domain.Order,
		ModulId:   domain.ModulId,
		CreatedId: domain.CreatedAt,
		UpdatedId: domain.UpdatedAt,
	}
}
