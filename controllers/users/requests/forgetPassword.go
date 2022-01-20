package requests

import "profcourse/business/users"

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

func (r ForgetPasswordRequest) ToDomain() *users.Domain {
	return &users.Domain{Email: r.Email}
}
