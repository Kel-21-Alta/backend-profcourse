package requests

import (
	"mime/multipart"
	"profcourse/business/courses"
)

type CreateCourseRequest struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	FileImage   *multipart.FileHeader `json:"file_image"`
}

func (r CreateCourseRequest) ToDomain() *courses.Domain {
	return &courses.Domain{
		Title:       r.Title,
		Description: r.Description,
		FileImage:   r.FileImage,
	}
}
