package requests

import "profcourse/business/users"

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  int8   `json:"role"`
}

func (u *UserRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}
}
