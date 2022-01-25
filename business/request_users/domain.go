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
	User User
	CreatedAt  time.Time
	UpdatedAt  time.Time

	RoleUser int8

	Query Query
}

type Query struct {
	Sort   string
	Search string
	Limit  int
	Offset int
}

type Category struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Name string
	Id string
}

type Usecase interface {
	CreateRequest(ctx context.Context, domain *Domain) (Domain, error)
	GetAllCategoryRequest(ctx context.Context) ([]Category, error)
	GetAllRequestUser(ctx context.Context, domain *Domain) ([]Domain, error)
	AdminGetAllRequestUser(ctx context.Context, domain *Domain) ([]Domain, error)
	DeleteRequestUser(ctx context.Context, domain *Domain) (Domain, error)
	UpdateRequestUser(ctx context.Context, domain *Domain) (Domain, error)
	GetOneRequestUser(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	CreateRequest(ctx context.Context, domain *Domain) (Domain, error)
	GetOneRequest(ctx context.Context, domain *Domain) (Domain, error)
	GetAllCategoryRequest(ctx context.Context) ([]Category, error)
	GetAllRequestUser(ctx context.Context, domain *Domain) ([]Domain, error)
	AdminGetAllRequestUser(ctx context.Context, domain *Domain) ([]Domain, error)
	DeleteRequestUser(ctx context.Context, domain *Domain) (Domain, error)
	UpdateRequestUser(ctx context.Context, domain *Domain) (Domain, error)
	GetOneRequestUser(ctx context.Context, domain *Domain) (Domain, error)
}
