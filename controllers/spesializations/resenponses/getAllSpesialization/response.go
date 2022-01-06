package getAllSpesialization

import (
	"profcourse/business/spesializations"
	"time"
)

type Response struct {
	Title     string    `json:"title"`
	UrlImage  string    `json:"url_image"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func FromDomain(domain *spesializations.Domain) Response {
	return Response{
		Title:     domain.Title,
		UrlImage:  domain.ImageUrl,
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
	}
}

func FromListDomain(domains []spesializations.Domain) []Response {
	var result []Response
	for _, domain := range domains {
		result = append(result, FromDomain(&domain))
	}
	return result
}
