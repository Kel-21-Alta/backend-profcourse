package updateCourse

import (
	"profcourse/business/courses"
	"time"
)

type UpdateCourseResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UrlImage    string    `json:"url_image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"`
}

func FromDomain(domain courses.Domain) *UpdateCourseResponse {
	var statusText string

	if int8(domain.Status) == 1 {
		statusText = "Publish"
	} else if int8(domain.Status) == 3 {
		statusText = "Pending"
	} else if int8(domain.Status) == 2 {
		statusText = "Draft"
	}

	return &UpdateCourseResponse{
		ID:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		UrlImage:    domain.ImgUrl,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		Status:      statusText,
	}
}
