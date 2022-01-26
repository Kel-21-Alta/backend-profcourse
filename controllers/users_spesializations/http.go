package users_spesializations

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/users_spesializations"
	controller "profcourse/controllers"
	"profcourse/controllers/users_spesializations/request"
	"profcourse/controllers/users_spesializations/responses/registerSpesialization"
)

type UsersSpesializationController struct {
	UsersSpesialization users_spesializations.Usecase
}

func NewUsersSpesializationController(userSpesialization users_spesializations.Usecase) *UsersSpesializationController {
	return &UsersSpesializationController{
		UsersSpesialization: userSpesialization,
	}
}

func (ctr *UsersSpesializationController) RegisterSpesialization(c echo.Context) error {
	var req request.RegisterSpesializationRequest

	err := c.Bind(&req)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	token, err := middlewares.ExtractClaims(c)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	var domain = req.ToDomain()

	domain.UserID = token.Userid

	ctx := c.Request().Context()
	result, err := ctr.UsersSpesialization.RegisterSpesialization(ctx, domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, registerSpesialization.FromDomain(result))
}