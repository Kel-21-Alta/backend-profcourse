package createSpesialization

import (
	"profcourse/business/spesializations"
)

type CreateSpesilizationRequest struct {
	FileImage   string   `json:"file_image"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Courses     []string `json:"courses"`
}

func (receiver CreateSpesilizationRequest) ToDomain() *spesializations.Domain {
	return &spesializations.Domain{
		Title:        receiver.Title,
		ImageUrl:    receiver.FileImage,
		Description: receiver.Description,
		Courses:     receiver.Courses,
	}
}
