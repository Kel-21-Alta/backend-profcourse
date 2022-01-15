package materies

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/business/materies"
	controller "profcourse/controllers"
	"profcourse/controllers/materies/requests"
	"profcourse/controllers/materies/responses/createMateries"
	"profcourse/controllers/materies/responses/updateMateri"
)

type MateriesController struct {
	MateriesUsecase materies.Usecase
}

func NewMateriesController(usecase materies.Usecase) *MateriesController {
	return &MateriesController{MateriesUsecase: usecase}
}

func (ctr MateriesController) CreateMateries(c echo.Context) error {
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

func (ctr MateriesController) UpdateMateri(c echo.Context) error {
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
