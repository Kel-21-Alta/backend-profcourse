package createSpesialization

import (
	"profcourse/business/spesializations"
	"time"
)

type CreateSpesializationResponse struct {
	ID          string    `json:"id"`
	UrlImage    string    `json:"url_image"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Courses     []string  `json:"courses"`
	CreatedAt   time.Time `json:"created_at"`
}

func FromDomain(domain *spesializations.Domain) CreateSpesializationResponse {
	return CreateSpesializationResponse{
		ID:          domain.ID,
		UrlImage:    domain.ImageUrl,
		Title:       domain.Title,
		Description: domain.Description,
		Courses:     domain.CourseIds,
		CreatedAt:   domain.CreatedAt,
	}
}
