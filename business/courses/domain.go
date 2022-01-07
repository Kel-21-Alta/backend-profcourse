package courses

import (
	"context"
	"time"
)

type InfoCurrentUser struct {
	CurrentUser string
	IsRegister  bool
	Progress    int
}

type Modul struct {
	NameModul string
	ModulID   string
}

type Rangking struct {
	UserId   string
	NameUser string
	Skor     int
}

type Domain struct {
	ID          string
	Title       string
	Description string
	ImgUrl      string
	TeacherId   string
	TeacherName string
	Status      int8
	StatusText  string

	CountCourse int

	CertificateId string

	// Info User yang saat ini login
	InfoUser InfoCurrentUser

	//Modul
	Moduls []Modul

	//Rangking/leaderboard
	Rangking []Rangking

	UserTakenCourse int

	// Params
	Limit         int
	SortBy        string
	Sort          string
	KeywordSearch string
	Offset        int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Summary struct {
	CountCourse int
}

type Usecase interface {
	CreateCourse(ctx context.Context, domain *Domain) (*Domain, error)
	GetOneCourse(ctx context.Context, domain *Domain) (*Domain, error)
	GetAllCourses(ctx context.Context, domain *Domain) (*[]Domain, error)
	GetCountCourse(ctx context.Context) (*Summary, error)
}

type Repository interface {
	CreateCourse(ctx context.Context, domain *Domain) (*Domain, error)
	GetOneCourse(ctx context.Context, domain *Domain) (*Domain, error)
	GetAllCourses(ctx context.Context, domain *Domain) (*[]Domain, error)
	GetCountCourse(ctx context.Context) (*Summary, error)
}
