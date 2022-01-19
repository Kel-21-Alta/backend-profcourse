package materies

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/materies"
	controller "profcourse/controllers"
	"profcourse/controllers/materies/requests"
	"profcourse/controllers/materies/responses/createMateries"
	"profcourse/controllers/materies/responses/getAllMateri"
	"profcourse/controllers/materies/responses/getOneMateri"
	"profcourse/controllers/materies/responses/updateMateri"
)

type MateriesController struct {
	MateriesUsecase materies.Usecase
}

func NewMateriesController(usecase materies.Usecase) *MateriesController {
	return &MateriesController{MateriesUsecase: usecase}
}

func (ctr *MateriesController) CreateMateries(c echo.Context) error {
	ctx := c.Request().Context()
	var req requests.CreateMateriesRequest

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	clean, err := ctr.MateriesUsecase.CreateMateri(ctx, req.ToDomain())

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusCreated, createMateries.FromDomain(clean))
}

func (ctr *MateriesController) UpdateMateri(c echo.Context) error {
	ctx := c.Request().Context()

	var req requests.UpdateMateriRequest

	req.ID = c.Param("materiid")

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	clean, err := ctr.MateriesUsecase.UpdateMateri(ctx, req.ToDomain())

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, updateMateri.FromDomain(clean))
}

func (ctr *MateriesController) DeleteMateries(c echo.Context) error {
	ctx := c.Request().Context()

	var domain materies.Domain
	domain.ID = c.Param("materiid")

	clean, err := ctr.MateriesUsecase.DeleteMateri(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	type Message struct {
		Message string
	}

	return controller.NewResponseSuccess(c, http.StatusOK, Message{Message: "Materi dengan id " + clean.ID + " telah dihapus"})
}

func (ctr *MateriesController) GetOneMateri(c echo.Context) error {
	ctx := c.Request().Context()

	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}
	var domain materies.Domain

	domain.ID = c.Param("materiid")
	domain.User.ID = token.Userid

	clean, err := ctr.MateriesUsecase.GetOneMateri(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getOneMateri.FormDomain(clean))
}

func (ctr *MateriesController) GetAllMateri(c echo.Context) error {
	var domain materies.Domain

	domain.ModulId = c.Param("modulid")
	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	domain.User.ID = token.Userid

	var ctx = c.Request().Context()
	clean, err := ctr.MateriesUsecase.GetAllMateri(ctx, &domain)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getAllMateri.FromDomain(clean))
}

func (ctr *MateriesController) UpdateProgressMateri(c echo.Context) error {

	var rec requests.ProgressMateriProgress

	if err := c.Bind(&rec); err != nil {
		return controller.NewResponseError(c, err)
	}

	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}
	rec.UserId = token.Userid

	ctx := c.Request().Context()
	_, err = ctr.MateriesUsecase.UpdateProgressMateri(ctx, rec.ToDomain())

	if err != nil {
		return controller.NewResponseError(c, err )
	}
	type Message struct {
		 Message string
	}
	return controller.NewResponseSuccess(c, http.StatusOK, Message{Message: "Success mengupdate progess"})
}
