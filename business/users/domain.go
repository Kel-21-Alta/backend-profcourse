package users

import (
	"context"
	"time"
)

type Domain struct {
	ID           string
	Name         string
	Email        string
	Password     string
	HashPassword string
	NoHp         string
	Birth        time.Time
	BirthPlace   string
	Bio          string
	ImgProfile   string
	Role         int8 // 1 for admin, 2 for user
	RoleText     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Token        string
}
type Usecase interface {
	CreateUser(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	CreateUser(ctx context.Context, domain Domain) (Domain, error)
	GetUserByEmail(ctx context.Context, email string) (Domain, error)
}
