package summary

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/business/summary"
	controller "profcourse/controllers"
	"profcourse/controllers/summary/responses/getAllSummary"
)

type SummaryController struct {
	SummaryUsecase summary.Usecase
}

func NewSummaryController(su summary.Usecase) *SummaryController {
	return &SummaryController{SummaryUsecase: su}
}

func (sc *SummaryController) GetAllSummary(c echo.Context) error {
	ctx := c.Request().Context()
	clean, err := sc.SummaryUsecase.GetAllSummary(ctx)
	if err != nil {
		return controller.NewResponseError(c, err)
	}
	return controller.NewResponseSuccess(c, http.StatusOK, getAllSummary.FromDomain(clean))
}
