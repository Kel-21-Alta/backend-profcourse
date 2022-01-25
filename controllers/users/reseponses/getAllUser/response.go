package getAllUser

import "profcourse/business/users"

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ImageUrl    string `json:"image_url"`
	TakenCourse int    `json:"taken_course"`
	Point       int    `json:"point"`
}

func FromListDomain(domain []users.Domain) []User {
	var list []User

	for _, user := range domain {
		list = append(list, User{
			Id:          user.ID,
			Name:        user.Name,
			ImageUrl:    user.ImgProfile,
			TakenCourse: user.TakenCourse,
			Point:       user.Point,
		})
	}
	return list
}
