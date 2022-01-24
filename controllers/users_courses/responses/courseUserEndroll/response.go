package courseUserEndroll

import "profcourse/business/users_courses"

type Courses struct {
	Title    string `json:"title"`
	CourseId string `json:"course_id"`
	Progress int    `json:"progress"`
	UrlImage string `json:"url_image"`
}

type CourseUserEndrollResponse struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	CountCourse int    `json:"count_course"`
	Courses     []Courses
}

func FromDomain(domain users_courses.User) *CourseUserEndrollResponse {
	var listCourse []Courses

	for _, course := range domain.Courses {
		listCourse = append(listCourse, Courses{
			Title: course.CourseTitle,
			CourseId: course.CourseId,
			Progress: course.Progres,
			UrlImage: course.UrlImage,
		})
	}

	return &CourseUserEndrollResponse{
		UserID:      domain.UserID,
		Name:        domain.Name,
		CountCourse: domain.CountCourse,
		Courses:     listCourse,
	}
}
