package getAllModul

import (
	"profcourse/business/moduls"
	"time"
)

type GetAllModulResponse struct {
	NameModul string `json:"name_modul"`
	ModulId   string `json:"modul_id"`
	Order     int    `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain moduls.Domain) GetAllModulResponse {
	return GetAllModulResponse{
		NameModul: domain.Title,
		ModulId:   domain.ID,
		Order:     domain.Order,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domains []moduls.Domain) []GetAllModulResponse {
	var listResponse []GetAllModulResponse

	for _, domain := range domains {
		listResponse = append(listResponse, FromDomain(domain))
	}

	return listResponse
}