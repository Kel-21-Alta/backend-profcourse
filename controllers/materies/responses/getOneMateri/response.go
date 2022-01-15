package getOneMateri

import (
	"profcourse/business/materies"
	"time"
)

type GetOneMateriResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	ModulId   string    `json:"modul_id"`
	Order     int    `json:"order"`
	UrlMateri string    `json:"url_materi"`
	Type      int       `json:"type"`
	TypeText  string    `json:"type_text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User CurrentUser `json:"user"`
}

type CurrentUser struct {
	ID          string
	CurrentTime string
	IsComplate  bool
}

func FormDomain(domain materies.Domain) *GetOneMateriResponse {
	return &GetOneMateriResponse{
		ID:        domain.ID,
		Title:     domain.Title,
		ModulId:   domain.ModulId,
		Order: 		int(domain.Order),
		UrlMateri: domain.UrlMateri,
		Type:      domain.Type,
		TypeText:  domain.TypeString,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		User:      CurrentUser(domain.User),
	}
}

