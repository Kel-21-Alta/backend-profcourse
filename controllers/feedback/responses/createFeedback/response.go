package createFeedback

import (
	"profcourse/business/feedback"
	"time"
)

type CreateFeedbackResponse struct {
	ID string `json:"id"`
	Rating float32 `json:"rating"`
	Review string `json:"review"`
	CourseId string `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
}

func FromDomain(domain feedback.Domain) *CreateFeedbackResponse {
	return &CreateFeedbackResponse{
		ID:        domain.ID,
		Rating:    domain.Rating,
		Review:    domain.Review,
		CourseId:  domain.CourseId,
		CreatedAt: domain.CreatedAt,
	}
}
