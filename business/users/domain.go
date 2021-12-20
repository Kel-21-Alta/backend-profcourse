package users

import "context"

type Domain struct {
	name string
}
type Usecase interface {
	TestClean(ctx context.Context, name string) (string, error)
}
