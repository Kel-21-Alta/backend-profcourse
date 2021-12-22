package requests

import "profcourse/business/users"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  int8   `json:"role"`
}

func (u *User) ToDomain() *users.Domain {
	return &users.Domain{
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}
}
