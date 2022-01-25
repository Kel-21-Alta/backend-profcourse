package users

import (
	"context"
	"time"
)

type Domain struct {
	ID           string
	Name         string
	Email        string
	Password     string
	HashPassword string
	NoHp         string
	Birth        time.Time

	BirthPlace string
	Bio        string
	ImgProfile string

	Role     int8 // 1 for admin, 2 for user
	RoleText string

	IdUser  string
	Message string

	TakenCourse int
	Point       int

	Query Query

	FileReport string

	CreatedAt   time.Time
	UpdatedAt   time.Time
	Token       string
	PasswordNew string
}

type Query struct {
	Sort   string
	SortBy string
	Limit  int
	Search string
	Offset int
}

type Summary struct {
	CountUser int
}

type Course struct {
	ID          string
	UserId      string
	CourseId    string
	Progres     int
	LastVideoId string
	LastModulId string
	Score       int

	CourseTitle string
	UrlImage string

	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	CreateUser(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
	LoginAdmin(ctx context.Context, domain Domain) (Domain, error)
	ForgetPassword(ctx context.Context, domain Domain) (Domain, error)
	GetCurrentUser(ctx context.Context, domain Domain) (Domain, error)
	ChangePassword(ctx context.Context, domain Domain) (Domain, error)
	DeleteUser(ctx context.Context, domain Domain) (Domain, error)
	GetCountUser(ctx context.Context) (*Summary, error)
	UpdateUser(ctx context.Context, domain Domain) (Domain, error)
	UpdateDataCurrentUser(ctx context.Context, domain *Domain) (Domain, error)
	GetAllUser(ctx context.Context, domain *Domain) ([]Domain, error)
	GenerateReportUser(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	CreateUser(ctx context.Context, domain Domain) (Domain, error)
	GetUserByEmail(ctx context.Context, email string) (Domain, error)
	UpdatePassword(ctx context.Context, domain Domain, hash string) (Domain, error)
	GetUserById(ctx context.Context, id string) (Domain, error)
	DeleteUser(ctx context.Context, domain Domain) (Domain, error)
	UpdateUser(ctx context.Context, domain Domain) (Domain, error)
	GetCountUser(ctx context.Context) (*Summary, error)
	UpdateDataCurrentUser(ctx context.Context, domain *Domain) (Domain, error)
	GetAllUser(ctx context.Context, domain *Domain) ([]Domain, error)
	GetCourseUser(ctx context.Context, domain *Domain) ([]Course, error)
}
