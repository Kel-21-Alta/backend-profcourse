package summary

import "golang.org/x/net/context"

type Domain struct {
	CountCourse int
	CountUser   int
	CountSpesialization int
}

type Usecase interface {
	GetSummary(ctx context.Context, domain *Domain)(*Domain, error)
}

type Repository interface {
	GetSummary(ctx context.Context, domain *Domain)(*Domain, error)
}
