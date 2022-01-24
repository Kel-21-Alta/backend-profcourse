package getAllCategoryRequestUser

import (
	"profcourse/business/request_users"
	"time"
)

type GetAllRequestUser struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromListDomain(domains []request_users.Category) []GetAllRequestUser {
	var list []GetAllRequestUser
	for _, domain := range domains {
		list = append(list, GetAllRequestUser{
			ID:        domain.ID,
			Title:     domain.Title,
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		})
	}

	return list
}
