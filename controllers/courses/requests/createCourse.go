package requests

import (
	"profcourse/business/courses"
)

type CreateCourseRequest struct {
	Title       string                `json:"title" form:"title"`
	Description string                `json:"description" form:"description"`
	FileImage   string 				`json:"file_image" form:"file_image"`
	UserId      string
}

func (r CreateCourseRequest) ToDomain() *courses.Domain {
	return &courses.Domain{
		Title:       r.Title,
		Description: r.Description,
		ImgUrl:   	r.FileImage,
		TeacherId:   r.UserId,
	}
}
