package reseponses

import (
	"profcourse/business/users"
	"time"
)

type UserCreated struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Role      int8      `json:"role"`
	RoleText  string    `json:"role_text"`
}

func FromDomain(domain users.Domain) UserCreated {
	return UserCreated{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		Role:      domain.Role,
		RoleText:  domain.RoleText,
	}
}
