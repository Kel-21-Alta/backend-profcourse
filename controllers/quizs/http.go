package quizs

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/business/quizs"
	controller "profcourse/controllers"
	"profcourse/controllers/quizs/requests"
	"profcourse/controllers/quizs/responses/createQuizs"
	"profcourse/controllers/quizs/responses/updateQuiz"
)

type QuizsController struct {
	QuizsUsecase quizs.Usecase
}

func NewQuizsController(usecase quizs.Usecase) *QuizsController {
	return &QuizsController{QuizsUsecase: usecase}
}

func (ctr *QuizsController) CreateQuiz(c echo.Context) error {
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

func (ctr *QuizsController) UpdateQuiz(c echo.Context) error {

	ctx := c.Request().Context()

	var req requests.UpdateQuizRequest

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}
	req.ID = c.Param("quizid")

	clean, err := ctr.QuizsUsecase.UpdateQuiz(ctx, req.ToDomain())

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, updateQuiz.FromDomain(clean))
}

func (ctr *QuizsController) DeleteQuiz(c echo.Context) error {
	var id string

	id = c.Param("quizid")
	ctx := c.Request().Context()
	resultId, err := ctr.QuizsUsecase.DeleteQuiz(ctx, id)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	type Message struct {
		Message string
	}
	return controller.NewResponseSuccess(c, http.StatusOK, Message{Message: "Quis dengan id " + resultId + " berhasil dihapus"})
}
