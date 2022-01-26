package users_spesializations

import (
	"golang.org/x/net/context"
	"time"
)

type Domain struct {
	ID string
	UserID string
	SpesializationID string
	Progress int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	RegisterSpesialization(ctx context.Context, domain *Domain)(Domain, error)
}

type Repository interface {
	RegisterSpesialization(ctx context.Context, domain *Domain)(Domain, error)
	GetEndRollSpesializationById(ctx context.Context, domain *Domain) (Domain, error)
}