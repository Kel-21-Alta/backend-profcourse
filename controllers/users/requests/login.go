package requests

import "profcourse/business/users"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *LoginRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    u.Email,
		Password: u.Password,
	}
}
