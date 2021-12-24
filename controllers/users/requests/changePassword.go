package requests

import "profcourse/business/users"

type ChangePasswordRequest struct {
	PasswordNew string `json:"password_new"`
	PasswordOld string `json:"password_old"`
	ID          string
}

func (r ChangePasswordRequest) ToDomain() users.Domain {
	return users.Domain{
		ID:          r.ID,
		Password:    r.PasswordOld,
		PasswordNew: r.PasswordNew,
	}
}

func (r ChangePasswordRequest) SetID(id string) ChangePasswordRequest {
	return ChangePasswordRequest{
		PasswordNew: r.PasswordNew,
		PasswordOld: r.PasswordOld,
		ID:          id,
	}
}
