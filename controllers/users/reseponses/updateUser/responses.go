package updateUser

import (
	"profcourse/business/users"
	"time"
)

type UserUpdated struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Role       int8      `json:"role"`
	RoleText   string    `json:"role_text"`
	NoHp       string    `json:"noHp"`
	Bio        string    `json:"bio"`
	Birth      time.Time `json:"birth"`
	BirthPlace string    `json:"birthplace"`
}

func FromDomain(domain users.Domain) UserUpdated {
	return UserUpdated{
		ID:         domain.ID,
		Name:       domain.Name,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		Role:       domain.Role,
		RoleText:   domain.RoleText,
		NoHp:       domain.NoHp,
		Bio:        domain.Bio,
		Birth:      domain.Birth,
		BirthPlace: domain.BirthPlace,
	}
}
