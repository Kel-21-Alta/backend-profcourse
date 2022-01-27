package requests

import "profcourse/business/courses"

type UpdateCourse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	FileImage   string `json:"file_image"`
	Status int8 `json:"status"`
}

func (c UpdateCourse) ToDomain() *courses.Domain {
	return &courses.Domain{Title: c.Title, Description: c.Description, ImgUrl: c.FileImage, Status: c.Status}
}
