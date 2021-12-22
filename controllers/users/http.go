package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/business/users"
	controller "profcourse/controllers"
	"profcourse/controllers/users/requests"
	"profcourse/controllers/users/reseponses"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{userUsecase: uc}
}

func (ctrl *UserController) CreateUser(c echo.Context) error {

	ctx := c.Request().Context()
	req := requests.User{}

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	clean, err := ctrl.userUsecase.CreateUser(ctx, *req.ToDomain())

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusCreated, reseponses.FromDomain(clean))
}
