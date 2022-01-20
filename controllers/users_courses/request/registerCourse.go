package request

import "profcourse/business/users_courses"

type RegisterCourseRequest struct {
	CourseId string `json:"course_id"`
}

func (r RegisterCourseRequest) ToDomain() *users_courses.Domain {
	return &users_courses.Domain{CourseId: r.CourseId}
}
