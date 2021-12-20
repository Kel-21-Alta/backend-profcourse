package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/business/users"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{userUsecase: uc}
}

func (ctrl *UserController) Test(c echo.Context) error {

	ctx := c.Request().Context()

	clean, err := ctrl.userUsecase.TestClean(ctx, "ikan")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, clean)
}
