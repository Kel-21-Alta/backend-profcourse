package getOneCourse

import "profcourse/business/courses"

type InfoCurrentUser struct {
	CurrentUser string `json:"current_user"`
	IsRegister  bool   `json:"is_register"`
	Progress    int    `json:"progress"`
}

type Modul struct {
	NameModul string `json:"name_modul"`
	ModulID   string `json:"modul_id"`
	Order     int    `json:"order"`
}

type Ranking struct {
	UserID   string `json:"user_id"`
	NameUser string `json:"name_user"`
	Skor     int    `json:"skor"`
}

type GetOneCourseResponses struct {
	InfoUser        InfoCurrentUser `json:"info_user"`
	CourseId        string          `json:"course_id"`
	NameCourse      string          `json:"name_course"`
	Description     string          `json:"description"`
	UrlImage        string          `json:"url_image"`
	Teacher         string          `json:"teacher"`
	TeacherId 		string `json:"teacher_id"`
	Moduls          []Modul         `json:"moduls"`
	UserTakenCourse int             `json:"user_taken_course"`
	Rangking        []Ranking       `json:"rangking"`
}

func FromDomain(domain *courses.Domain) *GetOneCourseResponses {

	var listModuls []Modul

	for _, modul := range domain.Moduls {
		listModuls = append(listModuls, Modul{
			NameModul: modul.NameModul,
			ModulID:   modul.ModulID,
			Order:     modul.Order,
		})
	}

	var listRangking []Ranking

	for _, rangking := range domain.Rangking {
		listRangking = append(listRangking, Ranking{
			UserID:   rangking.UserId,
			NameUser: rangking.NameUser,
			Skor:     rangking.Skor,
		})
	}

	return &GetOneCourseResponses{
		CourseId:    domain.ID,
		NameCourse:  domain.Title,
		Description: domain.Description,
		UrlImage:    domain.ImgUrl,
		Teacher:     domain.TeacherName,
		TeacherId: 		domain.TeacherId,
		InfoUser: InfoCurrentUser{
			CurrentUser: domain.InfoUser.CurrentUser,
			IsRegister:  domain.InfoUser.IsRegister,
			Progress:    domain.InfoUser.Progress,
		},
		Moduls:          listModuls,
		UserTakenCourse: domain.UserTakenCourse,
		Rangking:        listRangking,
	}
}
