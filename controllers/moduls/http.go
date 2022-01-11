package moduls

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/moduls"
	controller "profcourse/controllers"
	"profcourse/controllers/moduls/request"
	"profcourse/controllers/moduls/responses/createModul"
	"profcourse/controllers/moduls/responses/getOneModul"
)

type ModulController struct {
	ModulsUsecase moduls.Usecase
}

func NewModulsController(usecase moduls.Usecase) *ModulController {
	return &ModulController{ModulsUsecase: usecase}
}

func (ctr ModulController) CreateModul(c echo.Context) error {
	//	mendapatkan data dari request
	req := request.CreateModulsRequest{}
	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}
	// Mendapatkan siapa user yang membuat modul
	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}
	domain := req.ToDomain()
	domain.UserId = token.Userid

	// Usecase
	ctx := c.Request().Context()
	clean, err := ctr.ModulsUsecase.CreateModul(ctx, domain)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusCreated, createModul.FromDomain(clean))
}

func (ctr ModulController) GetOneModul(c echo.Context) error {

	ctx := c.Request().Context()
	clean, err := ctr.ModulsUsecase.GetOneModul(ctx, &moduls.Domain{ID: c.Param("modulid")})

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getOneModul.FromDomain(&clean))
}
