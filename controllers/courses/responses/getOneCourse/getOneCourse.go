package getOneCourse

import "profcourse/business/courses"

type GetOneCourseResponses struct {
	CourseId    string `json:"course_id"`
	NameCourse  string `json:"name_course"`
	Description string `json:"description"`
	UrlImage string `json:"url_image"`
	Teacher string `json:"teacher"`
}

func FromDomain(domain *courses.Domain) *GetOneCourseResponses {
	return &GetOneCourseResponses{
		CourseId:    domain.ID,
		NameCourse:  domain.Title,
		Description: domain.Description,
		UrlImage:    domain.ImgUrl,
		Teacher:     domain.TeacherName,
	}
}
