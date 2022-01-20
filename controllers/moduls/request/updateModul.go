package request

import (
	"profcourse/business/moduls"
)

type UpdateModulRequest struct {
	CourseId string `json:"course_id"`
	Title    string `json:"title"`
	Order int `json:"order"`
}

func (r UpdateModulRequest) ToDomain() *moduls.Domain {
	return &moduls.Domain{
		Title:         r.Title,
		Order:         r.Order,
		CourseId:      r.CourseId,
	}
}
