package getOneModul

import (
	"profcourse/business/moduls"
	"time"
)

type GetOneModulResponse struct {
	ID        string    `json:"id"`
	CourseId  string    `json:"course_id"`
	Title     string    `json:"title"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	JumlahMateri int `json:"jumlah_materi"`
	Materi []Materi `json:"materi"`
}

type Materi struct {
	UrlMateri   string `json:"url_materi"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Order       int8 `json:"order"`
	IsComplate  bool `json:"is_complate"`
	CurrentTime string `json:"current_time"`
}

func FromDomain(domain *moduls.Domain) *GetOneModulResponse {
	var listMateri []Materi
	for _, materi := range domain.Materi {
		listMateri = append(listMateri, Materi{
			UrlMateri:   materi.UrlMateri,
			Type:        materi.Type,
			Title:       materi.Title,
			Order:       materi.Order,
			IsComplate:  materi.IsComplate,
			CurrentTime: materi.CurrentTime,
		})
	}
	return &GetOneModulResponse{
		ID:        domain.ID,
		CourseId:  domain.CourseId,
		Title:     domain.Title,
		Order:     domain.Order,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		JumlahMateri: domain.JumlahMateri,
		Materi: listMateri,
	}
}
