package spesializations

import (
	"golang.org/x/net/context"
	"time"
)

type Domain struct {
	ID            string
	Title         string
	ImageUrl      string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CertificateId string
	Courses       []string
	MakerRole     int

	//Params
	Limit         int
	SortBy        string
	Sort          string
	KeywordSearch string
	Offset        int
}

type Repository interface {
	CreateSpasialization(ctx context.Context, domain *Domain) (Domain, error)
	GetAllSpesializations(ctx context.Context, domain *Domain) ([]Domain, error)
}

type Usecase interface {
	CreateSpasialization(ctx context.Context, domain *Domain) (Domain, error)
	GetAllSpesializations(ctx context.Context, domain *Domain) ([]Domain, error)
}
