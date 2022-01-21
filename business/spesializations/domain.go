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
	CourseIds     []string
	Courses       []Course
	MakerRole     int

	//Params
	Limit         int
	SortBy        string
	Sort          string
	KeywordSearch string
	Offset        int
}

type Course struct {
	ID          string
	Title       string
	Rating      float32
	Description string
}

type Summary struct {
	CountSpesialization int
}

type Repository interface {
	CreateSpasialization(ctx context.Context, domain *Domain) (Domain, error)
	GetOneSpesialization(ctx context.Context, domain *Domain) (Domain, error)
	GetAllSpesializations(ctx context.Context, domain *Domain) ([]Domain, error)
	GetCountSpesializations(ctx context.Context) (Summary, error)
}

type Usecase interface {
	GetOneSpesialization(ctx context.Context, domain *Domain) (Domain, error)
	CreateSpasialization(ctx context.Context, domain *Domain) (Domain, error)
	GetAllSpesializations(ctx context.Context, domain *Domain) ([]Domain, error)
	GetCountSpesializations(ctx context.Context) (Summary, error)
}
