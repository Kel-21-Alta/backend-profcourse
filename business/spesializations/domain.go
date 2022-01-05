package spesializations

import (
	"golang.org/x/net/context"
	"time"
)

type Domain struct {
	ID            string
	Name          string
	ImageUrl      string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CertificateId string
	Courses       []string
	MakerRole     int
}

type Repository interface {
	CreateSpasialization(ctx context.Context, domain *Domain) (*Domain, error)
}

type Usecase interface {
	CreateSpasialization(ctx context.Context, domain *Domain) (*Domain, error)
}
