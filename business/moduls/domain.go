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

	UserId           string
	UserIdMakeCourse string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Usecase interface {
	CreateModul(ctx context.Context, domain *Domain) (*Domain, error)
}

type Repository interface {
	CreateModul(ctx context.Context, domain *Domain) (*Domain, error)
}
