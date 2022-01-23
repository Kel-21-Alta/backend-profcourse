package users_courses

import (
	"context"
	"time"
)

type Domain struct {
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

type User struct {
	UserID      string
	Name        string
	CountCourse int
	Courses 	[]Domain
}

type Usecase interface {
	UserRegisterCourse(ctx context.Context, domain *Domain) (*Domain, error)
	UpdateProgressCourse(ctx context.Context, domain *Domain) (Domain, error)
	GetOneUserCourse(ctx context.Context, domain *Domain) (Domain, error)
	UpdateScoreCourse(ctx context.Context, domain *Domain) (Domain, error)
	GetUserCourseEndroll(ctx context.Context, domain *User) (User, error)
}

type Repository interface {
	UserRegisterCourse(ctx context.Context, domain *Domain) (*Domain, error)
	GetEndRollCourseUserById(ctx context.Context, domain *Domain) (*Domain, error)
	UpdateProgressCourse(ctx context.Context, domain *Domain) (Domain, error)
	GetOneUserCourse(ctx context.Context, domain *Domain) (Domain, error)
	UpdateScoreCourse(ctx context.Context, domain *Domain) (Domain, error)
	GetUserCourseEndroll(ctx context.Context, domain *User) (User, error)
}
