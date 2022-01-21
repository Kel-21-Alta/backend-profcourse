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
	Score int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	UserRegisterCourse(ctx context.Context, domain *Domain) (*Domain, error)
	UpdateProgressCourse(ctx context.Context, domain *Domain) (Domain, error)
	GetOneUserCourse(ctx context.Context, domain *Domain) (Domain, error)
	UpdateScoreCourse(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	UserRegisterCourse(ctx context.Context, domain *Domain) (*Domain, error)
	GetEndRollCourseUserById(ctx context.Context, domain *Domain) (*Domain, error)
	UpdateProgressCourse(ctx context.Context, domain *Domain) (Domain, error)
	GetOneUserCourse(ctx context.Context, domain *Domain) (Domain, error)
	UpdateScoreCourse(ctx context.Context, domain *Domain) (Domain, error)
}
