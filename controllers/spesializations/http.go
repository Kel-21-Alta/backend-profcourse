package spesializations

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/spesializations"
	controller "profcourse/controllers"
	_requestCreateSpesialization "profcourse/controllers/spesializations/requests/createSpesialization"
	_responseCreateSpesialization "profcourse/controllers/spesializations/resenponses/createSpesialization"
	"profcourse/controllers/spesializations/resenponses/getAllSpesialization"
	"profcourse/controllers/spesializations/resenponses/getOneSpesialization"
	"strconv"
)

type SpesializationController struct {
	SpesializationUsecase spesializations.Usecase
}

func NewSpesializationController(usecase spesializations.Usecase) *SpesializationController {
	return &SpesializationController{SpesializationUsecase: usecase}
}

func (sp *SpesializationController) GetAllSpesialization(c echo.Context) error {
	ctx := c.Request().Context()
	var domain spesializations.Domain
	var err error

	domain.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	domain.Offset, _ = strconv.Atoi(c.QueryParam("offset"))
	domain.Sort = c.QueryParam("sort")
	domain.SortBy = c.QueryParam("sortby")
	domain.KeywordSearch = c.QueryParam("s")

	// Usecase
	clean, err := sp.SpesializationUsecase.GetAllSpesializations(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	if len(clean) <= 0 {
		return controller.NewResponseSuccess(c, http.StatusOK, getAllSpesialization.ResponseMessage{Message: "Data Tidak ada"})
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getAllSpesialization.FromListDomain(clean))
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
	return controller.NewResponseSuccess(c, http.StatusCreated, _responseCreateSpesialization.FromDomain(&clean))
}

func (sp *SpesializationController) GetOneSpesialization(c echo.Context) error {
	ctx := c.Request().Context()
	var domain *spesializations.Domain

	domain.ID = c.Param("spesializationid")
	clean, err := sp.SpesializationUsecase.GetOneSpesialization(ctx, domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getOneSpesialization.FromDomain(clean))
}
