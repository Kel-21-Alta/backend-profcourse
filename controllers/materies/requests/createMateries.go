package requests

import (
	"profcourse/business/materies"
)

type CreateMateriesRequest struct {
	ModulId    string `json:"modul_id"`
	TypeMateri int `json:"type_materi"`
	Title      string `json:"title"`
	FileMateri string `json:"file_materi"`
	Order      int    `json:"order"`
}

func (receiver CreateMateriesRequest) ToDomain() *materies.Domain {
	return &materies.Domain{
		Title:     receiver.Title,
		ModulId:   receiver.ModulId,
		Order:     receiver.Order,
		Type:      receiver.TypeMateri,
		UrlMateri: receiver.FileMateri,
	}
}
