package getAllCourses

import (
	"profcourse/business/courses"
	"time"
)

type GetAllCoursesResponse struct {
	Title    string `json:"title"`
	UrlImage string `json:"url_image"`
	CourseId string `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain *courses.Domain) *GetAllCoursesResponse {
	return &GetAllCoursesResponse{
		Title:    domain.Title,
		UrlImage: domain.ImgUrl,
		CourseId: domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domains *[]courses.Domain) *[]GetAllCoursesResponse {
	var getAllCourses []GetAllCoursesResponse
	for _, domain := range *domains {
		getAllCourses = append(getAllCourses, *FromDomain(&domain))
	}
	return &getAllCourses
}
