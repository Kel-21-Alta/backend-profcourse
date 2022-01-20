package requests

import (
	"profcourse/business/materies"
)

type ProgressMateriProgress struct {
	MateriID   string `json:"materi_id"`
	IsComplate bool `json:"is_complate"`
	CurrentTime string `json:"current_time"`
	UserId string `json:"user_id"`
	CourseId string `json:"course_id"`
}

func (p ProgressMateriProgress) ToDomain() *materies.Domain {
	return &materies.Domain{
		ID:         p.MateriID,
		User:       materies.CurrentUser{
			ID:          p.UserId,
			CurrentTime: p.CurrentTime,
			IsComplate:  p.IsComplate,
			CourseId: p.CourseId,
		},
	}
}
