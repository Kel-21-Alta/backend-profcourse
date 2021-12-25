package createCourse

import (
	"profcourse/business/courses"
	"time"
)

type CreateCourseResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UrlImage    string    `json:"url_image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"`
}

func FromDomain(domain *courses.Domain) *CreateCourseResponse {
	return &CreateCourseResponse{
		ID:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		UrlImage:    domain.ImgUrl,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		Status:      domain.StatusText,
	}
}
