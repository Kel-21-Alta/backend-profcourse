package getOneSpesialization

import "profcourse/business/spesializations"

type Course struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
}

type Response struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageUrl    string   `json:"image_url"`
	Courses     []Course `json:"courses"`
}

func FromDomain(domain spesializations.Domain) Response {
	var listCourse []Course
	for _, course := range domain.Courses {
		listCourse = append(listCourse, Course{
			Title:       course.Title,
			Description: course.Description,
			Rating:      course.Rating,
		})
	}

	return Response{
		ID:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		ImageUrl:    domain.ImageUrl,
		Courses:     listCourse,
	}
}
