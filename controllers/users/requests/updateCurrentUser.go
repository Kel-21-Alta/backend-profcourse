package requests

import (
	"profcourse/business/users"
	"time"
)

type UpdateCurrentUser struct {
	Name       string `json:"name"`
	NoHp       string `json:"noHp"`
	Bio        string `json:"bio"`
	Birth      string `json:"birth"`
	BirthPlace string `json:"birthplace"`
	Profile    string `json:"profile"`
}

func (u *UpdateCurrentUser) ToDomain() *users.Domain {
	birth, _ := time.Parse("2006-01-02", u.Birth)
	return &users.Domain{
		ImgProfile: u.Profile,
		Name:       u.Name,
		NoHp:       u.NoHp,
		Bio:        u.Bio,
		Birth:      birth,
		BirthPlace: u.BirthPlace,
	}
}
