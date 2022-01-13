package requests

import (
	"profcourse/business/users"
	"time"
)

type UpdateCurrentUser struct {
	Name       string    `json:"name"`
	NoHp       string    `json:"noHp"`
	Bio        string    `json:"bio"`
	Birth      time.Time `json:"birth"`
	BirthPlace string    `json:"birthplace"`
	Profile    string    `json:"profile"`
}

func (u *UpdateCurrentUser) ToDomain() *users.Domain {
	return &users.Domain{
		ImgProfile: u.Profile,
		Name:       u.Name,
		NoHp:       u.NoHp,
		Bio:        u.Bio,
		Birth:      u.Birth,
		BirthPlace: u.BirthPlace,
	}
}
