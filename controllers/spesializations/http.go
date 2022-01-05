package spesializations

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/spesializations"
	controller "profcourse/controllers"
	_requestCreateSpesialization "profcourse/controllers/spesializations/requests/createSpesialization"
	_responseCreateSpesialization "profcourse/controllers/spesializations/resenponses/createSpesialization"
)

type SpesializationController struct {
	SpesializationUsecase spesializations.Usecase
}

func NewSpesializationController(usecase spesializations.Usecase) *SpesializationController {
	return &SpesializationController{SpesializationUsecase: usecase}
}

func (sp *SpesializationController) CreateSpesialization(c echo.Context) error {
	var req _requestCreateSpesialization.CreateSpesilizationRequest

	token, err := middlewares.ExtractClaims(c)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}
	domain := req.ToDomain()
	domain.MakerRole = int(token.Role)
	ctx := c.Request().Context()
	clean, err := sp.SpesializationUsecase.CreateSpasialization(ctx, domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}
	return controller.NewResponseSuccess(c, http.StatusCreated, _responseCreateSpesialization.FromDomain(clean))
}
