package request

import (
	"profcourse/business/moduls"
)

type CreateModulsRequest struct {
	CourseId string `json:"course_id"`
	Title    string `json:"title"`
	Order    int    `json:"order"`
}

func (receiver CreateModulsRequest) ToDomain() *moduls.Domain {
	return &moduls.Domain{
		Title:    receiver.Title,
		Order:    receiver.Order,
		CourseId: receiver.CourseId,
	}
}
