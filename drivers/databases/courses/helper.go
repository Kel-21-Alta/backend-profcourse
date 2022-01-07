package courses

import (
	"profcourse/business/courses"
)

type RegisteredUser struct {
	UserId   string
	NameUser string
	Skor     int
	Progress int
}

type InfoCurrentUser struct {
	IsRegister bool
	Progress   int
}

func FromRegiteredUserToDomain(domain *courses.Domain, users []RegisteredUser) *courses.Domain {

	var listRangking []courses.Rangking
	var infoCurrentUser InfoCurrentUser

	for _, user := range users {
		if user.UserId == domain.InfoUser.CurrentUser {
			infoCurrentUser.IsRegister = true
			infoCurrentUser.Progress = user.Progress
		}
		listRangking = append(listRangking, courses.Rangking{
			UserId:   user.UserId,
			NameUser: user.NameUser,
			Skor:     user.Skor,
		})
	}

	return &courses.Domain{
		ID:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		ImgUrl:      domain.ImgUrl,
		TeacherId:   domain.TeacherId,
		TeacherName: domain.TeacherName,
		Status:      domain.Status,
		StatusText:  domain.StatusText,
		InfoUser: courses.InfoCurrentUser{
			CurrentUser: domain.InfoUser.CurrentUser,
			IsRegister:  infoCurrentUser.IsRegister,
			Progress:    infoCurrentUser.Progress,
		},
		Moduls:          domain.Moduls,
		Rangking:        listRangking,
		UserTakenCourse: domain.UserTakenCourse,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
	}
}
