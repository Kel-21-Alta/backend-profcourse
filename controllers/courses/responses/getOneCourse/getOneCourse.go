package getOneCourse

import "profcourse/business/courses"

type InfoCurrentUser struct {
	CurrentUser string `json:"current_user"`
	IsRegister bool `json:"is_register"`
	Progress int `json:"progress"`
}

type GetOneCourseResponses struct {
	CourseId    string `json:"course_id"`
	NameCourse  string `json:"name_course"`
	Description string `json:"description"`
	UrlImage string `json:"url_image"`
	Teacher string `json:"teacher"`
	InfoUser InfoCurrentUser
}

func FromDomain(domain *courses.Domain) *GetOneCourseResponses {
	return &GetOneCourseResponses{
		CourseId:    domain.ID,
		NameCourse:  domain.Title,
		Description: domain.Description,
		UrlImage:    domain.ImgUrl,
		Teacher:     domain.TeacherName,
		InfoUser: InfoCurrentUser{
			CurrentUser: domain.InfoUser.CurrentUser,
			IsRegister:  domain.InfoUser.IsRegister,
			Progress:    domain.InfoUser.Progress,
		},
	}
}
