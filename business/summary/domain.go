package summary

import "golang.org/x/net/context"

type Domain struct {
	CountCourse         int
	CountUser           int
	CountSpesialization int
}

type Usecase interface {
	GetAllSummary(ctx context.Context) (*Domain, error)
}
