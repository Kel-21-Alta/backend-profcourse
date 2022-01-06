package createModul

import (
	"profcourse/business/moduls"
	"time"
)

type CreateModulResponse struct {
	ID        string    `json:"id"`
	CourseId  string    `json:"course_id"`
	Title     string    `json:"title"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain *moduls.Domain) *CreateModulResponse {
	return &CreateModulResponse{
		ID:        domain.ID,
		CourseId:  domain.CourseId,
		Title:     domain.Title,
		Order:     domain.Order,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
