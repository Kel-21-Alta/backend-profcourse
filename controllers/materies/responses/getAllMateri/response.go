package getAllMateri

import (
	"profcourse/business/materies"
)

type GetAllMateriResponse struct {
	JumlahMateri int      `json:"jumlah_materi"`
	Materi       []Materi `json:"materi"`
}

type Materi struct {
	ID string `json:"id"`
	UrlMateri   string `json:"url_materi"`
	Type        int    `json:"type"`
	TypeString  string `json:"type_string"`
	Title       string `json:"title"`
	Order       int8   `json:"order"`
	IsComplate  bool   `json:"is_complate"`
	CurrentTime string `json:"current_time"`
}

func FromDomain(domain materies.AllMateriModul) *GetAllMateriResponse {
	var listMateri []Materi
	for _, materi := range domain.Materi {
		listMateri = append(listMateri, Materi{
			ID: materi.ID,
			UrlMateri:   materi.UrlMateri,
			Type:        materi.Type,
			TypeString:  materi.TypeString,
			Title:       materi.Title,
			Order: 		int8(materi.Order),
			IsComplate:  materi.User.IsComplate,
			CurrentTime: materi.User.CurrentTime,
		})
	}
	
	return &GetAllMateriResponse{
		JumlahMateri: domain.JawabanMateri,
		Materi:       listMateri,
	}
}
