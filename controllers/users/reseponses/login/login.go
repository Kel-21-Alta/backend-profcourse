package login

import "profcourse/business/users"

type LoginResponses struct {
	Token    string `json:"token"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Profile  string `json:"profile"`
	Role     int8   `json:"role"`
	RoleText string `json:"role_text"`
}

func FromDomain(domain users.Domain) *LoginResponses {
	return &LoginResponses{
		Token:    domain.Token,
		Id:       domain.ID,
		Name:     domain.Name,
		Email:    domain.Email,
		Profile:  domain.ImgProfile,
		Role:     domain.Role,
		RoleText: domain.RoleText,
	}
}
