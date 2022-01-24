package requests

import (
	"profcourse/business/feedback"
)

type CreateFeedbackRequest struct {
	Rating   float32 `json:"rating"`
	Review   string  `json:"review"`
	CourseId string  `json:"course_id"`
}

func (receiver *CreateFeedbackRequest) Todomain() *feedback.Domain {
	return &feedback.Domain{
		CourseId: receiver.CourseId,
		Review:   receiver.Review,
		Rating:   receiver.Rating,
	}
}
