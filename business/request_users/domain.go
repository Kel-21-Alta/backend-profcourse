package request_users

import (
	"context"
	"time"
)

type Domain struct {
	Id         string
	UserId     string
	CategoryID string
	Request    string
	Category   Category
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Category struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	CreateRequest(ctx context.Context, domain *Domain) (Domain, error)
	GetAllCategoryRequest(ctx context.Context) ([]Category, error)
}

type Repository interface {
	CreateRequest(ctx context.Context, domain *Domain) (Domain, error)
	GetOneRequest(ctx context.Context, domain *Domain) (Domain, error)
	GetAllCategoryRequest(ctx context.Context) ([]Category, error)
}
