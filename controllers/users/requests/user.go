package requests

import "profcourse/business/users"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) ToDomain() *users.Domain {
	return &users.Domain{
		Name:  u.Name,
		Email: u.Email,
	}
}
