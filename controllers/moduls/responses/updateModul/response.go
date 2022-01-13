package updateModul

import (
	"profcourse/business/moduls"
	"time"
)

type Response struct {
	ID        string    `json:"id"`
	CourseId  string    `json:"course_id"`
	Title     string    `json:"title"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain moduls.Domain) Response {
	return Response{
		ID:        domain.ID,
		CourseId:  domain.CourseId,
		Title:     domain.Title,
		Order:     domain.Order,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}