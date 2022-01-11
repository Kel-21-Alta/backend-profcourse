package moduls

import (
	"context"
	"time"
)

type Domain struct {
	ID       string
	Title    string
	Order    int
	CourseId string

	Materi       []Materi
	JumlahMateri int

	UserMakeModul string
	RoleUser      int8
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Materi struct {
	UrlMateri   string
	Type        string
	Title string
	Order       int8
	IsComplate  bool
	CurrentTime string
}

type Usecase interface {
	CreateModul(ctx context.Context, domain *Domain) (Domain, error)
	GetOneModul(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	CreateModul(ctx context.Context, domain *Domain) (Domain, error)
	GetOneModul(ctx context.Context, domain *Domain) (Domain, error)
}
