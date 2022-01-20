package requests

import "profcourse/business/materies"

type UpdateMateriRequest struct {
	ID string `json:"id"`
	ModulId    string `json:"modul_id"`
	TypeMateri int `json:"type_materi"`
	Title      string `json:"title"`
	FileMateri string `json:"file_materi"`
	Order      int    `json:"order"`
}

func (receiver *UpdateMateriRequest) ToDomain() *materies.Domain {
	return &materies.Domain{
		ID: receiver.ID,
		Title:     receiver.Title,
		ModulId:   receiver.ModulId,
		Order:     receiver.Order,
		Type:      receiver.TypeMateri,
		UrlMateri: receiver.FileMateri,
	}
}
