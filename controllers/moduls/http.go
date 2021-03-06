package moduls

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/moduls"
	controller "profcourse/controllers"
	"profcourse/controllers/moduls/request"
	"profcourse/controllers/moduls/responses/createModul"
	"profcourse/controllers/moduls/responses/getAllModul"
	"profcourse/controllers/moduls/responses/getOneModul"
	"profcourse/controllers/moduls/responses/updateModul"
)

type ModulController struct {
	ModulsUsecase moduls.Usecase
}

func NewModulsController(usecase moduls.Usecase) *ModulController {
	return &ModulController{ModulsUsecase: usecase}
}

func (ctr *ModulController) CreateModul(c echo.Context) error {
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
	domain.UserMakeModul = token.Userid
	domain.RoleUser = token.Role

	// Usecase
	ctx := c.Request().Context()
	clean, err := ctr.ModulsUsecase.CreateModul(ctx, domain)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusCreated, createModul.FromDomain(&clean))
}

func (ctr *ModulController) GetOneModul(c echo.Context) error {

	ctx := c.Request().Context()
	clean, err := ctr.ModulsUsecase.GetOneModul(ctx, &moduls.Domain{ID: c.Param("modulid")})

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getOneModul.FromDomain(&clean))
}

func (ctr *ModulController) UpdateModul(c echo.Context) error {

	ctx := c.Request().Context()

	var req request.UpdateModulRequest
	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	token, err := middlewares.ExtractClaims(c)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	domain := req.ToDomain()
	domain.ID = c.Param("modulid")
	domain.UserMakeModul = token.Userid
	domain.RoleUser = token.Role

	clean, err := ctr.ModulsUsecase.UpdateModul(ctx, domain)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, updateModul.FromDomain(clean))
}

func (ctr *ModulController) DeleteModul(c echo.Context) error {
	var domain moduls.Domain

	token, err := middlewares.ExtractClaims(c)

	domain.ID = c.Param("modulid")
	domain.UserMakeModul = token.Userid
	domain.RoleUser = token.Role

	ctx := c.Request().Context()
	clean, err := ctr.ModulsUsecase.DeleteModul(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	type Message struct {
		Message string `json:"message"`
	}

	return controller.NewResponseSuccess(c, http.StatusOK, Message{Message: string(clean)})
}

func (ctr *ModulController) GetAllModulCourse(c echo.Context) error {
	ctx := c.Request().Context()

	var domain moduls.Domain

	domain.CourseId = c.Param("courseid")

	clean, err := ctr.ModulsUsecase.GetAllModulCourse(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getAllModul.FromListDomain(clean))
}
