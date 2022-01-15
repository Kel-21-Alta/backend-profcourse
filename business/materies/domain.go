package materies

import (
	"golang.org/x/net/context"
	"time"
)

type Domain struct {
	ID         string
	Title      string
	ModulId    string
	Order      int
	Type       int
	TypeString string
	UrlMateri  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	CreateMateri(ctx context.Context, domain *Domain) (Domain, error)
	DeleteMateri(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	DeleteMateri(ctx context.Context, domain *Domain) (Domain, error)
	CreateMateri(ctx context.Context, domain *Domain) (Domain, error)
}
