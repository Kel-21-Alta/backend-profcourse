package requests

import (
	"profcourse/business/users"
	"time"
)

type UpdateUserRequest struct {
	IdUser     string    `json:"user_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Role       int8      `json:"role"`
	NoHp       string    `json:"noHp"`
	Bio        string    `json:"bio"`
	Birth      time.Time `json:"birth"`
	BirthPlace string    `json:"birthplace"`
}

func (u *UpdateUserRequest) ToDomain() *users.Domain {
	return &users.Domain{
		IdUser:     u.IdUser,
		Role:       u.Role,
		Email:      u.Email,
		Name:       u.Name,
		NoHp:       u.NoHp,
		Bio:        u.Bio,
		Birth:      u.Birth,
		BirthPlace: u.BirthPlace,
	}
}
