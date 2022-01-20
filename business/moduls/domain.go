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

	Course Course

	UserMakeModul string
	RoleUser      int8
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Course struct {
	ID        string
	TeacherId string
}

type Materi struct {
	UrlMateri   string
	Type        int
	TypeString  string
	Title       string
	Order       int8
	IsComplate  bool
	CurrentTime string
}

type Message string

type Usecase interface {
	CreateModul(ctx context.Context, domain *Domain) (Domain, error)
	GetOneModul(ctx context.Context, domain *Domain) (Domain, error)
	UpdateModul(ctx context.Context, domain *Domain) (Domain, error)
	DeleteModul(ctx context.Context, domain *Domain) (Message, error)
	GetAllModulCourse(ctx context.Context, domain *Domain) ([]Domain, error)
}

type Repository interface {
	CreateModul(ctx context.Context, domain *Domain) (Domain, error)
	GetOneModul(ctx context.Context, domain *Domain) (Domain, error)
	UpdateModul(ctx context.Context, domain *Domain) (Domain, error)
	DeleteModul(ctx context.Context, id string) (Message, error)
	GetOneModulWithCourse(ctx context.Context, domain *Domain) (Domain, error)
	GetAllModulCourse(ctx context.Context, domain *Domain) ([]Domain, error)
}
