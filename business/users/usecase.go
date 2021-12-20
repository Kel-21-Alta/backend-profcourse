package users

import (
	"context"
	"fmt"
	"time"
)

type userUsecase struct {
	contextTimeout time.Duration
}

func (u userUsecase) TestClean(ctx context.Context, name string) (string, error) {
	fmt.Println("implement me", name)
	return "hai " + name, nil
}

func NewUserUsecase() Usecase {
	return &userUsecase{}
}
