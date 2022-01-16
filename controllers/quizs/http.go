package quizs

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/business/quizs"
	controller "profcourse/controllers"
	"profcourse/controllers/quizs/requests"
	"profcourse/controllers/quizs/responses/createQuizs"
)

type QuizsController struct {
	QuizsUsecase quizs.Usecase
}

func NewQuizsController(usecase quizs.Usecase) *QuizsController {
	return &QuizsController{QuizsUsecase: usecase}
}

func (ctr QuizsController) CreateQuiz(c echo.Context) error {
	ctx := c.Request().Context()

	var req requests.CreateQuizRequest
	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	clean, err := ctr.QuizsUsecase.CreateQuiz(ctx, req.ToDomain())

	if err != nil {
		return controller.NewResponseError(c, err)
	}
	return controller.NewResponseSuccess(c, http.StatusOK, createQuizs.FromDomain(clean))
}
