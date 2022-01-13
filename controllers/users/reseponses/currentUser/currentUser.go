package currentUser

import (
	"profcourse/business/users"
	"time"
)

type UserCreated struct {
	ID         string `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	RoleText   string `json:"role"`
	Bio        string `json:"bio"`
	UrlImage   string `json:"url_image"`
	Birth      string `json:"birth"`
	NoHp       string `json:"no_hp"`
	BirthPlace string `json:"birth_place"`
}

func FromDomain(domain users.Domain) UserCreated {
	return UserCreated{
		ID:         domain.ID,
		Name:       domain.Name,
		Email:      domain.Email,
		RoleText:   domain.RoleText,
		Bio:        domain.Bio,
		UrlImage:   domain.ImgProfile,
		Birth:      domain.Birth.Format(time.RFC822),
		NoHp:       domain.NoHp,
		BirthPlace: domain.BirthPlace,
	}
}
