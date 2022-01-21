package quizs

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/quizs"
	controller "profcourse/controllers"
	"profcourse/controllers/quizs/requests"
	"profcourse/controllers/quizs/responses/createQuizs"
	"profcourse/controllers/quizs/responses/getAllQuizModul"
	"profcourse/controllers/quizs/responses/getOneQuiz"
	"profcourse/controllers/quizs/responses/jawabanQuiz"
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

func (ctr *QuizsController) GetAllQuizModul(c echo.Context) error {

	ctx := c.Request().Context()

	var domain quizs.Domain

	domain.ModulId = c.Param("modulid")

	clean, err := ctr.QuizsUsecase.GetAllQuizModul(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getAllQuizModul.FromListDomain(clean))
}

func (ctr *QuizsController) GetOneQuiz(c echo.Context) error {
	ctx := c.Request().Context()
	var domain quizs.Domain
	domain.ID = c.Param("quizid")

	clean, err := ctr.QuizsUsecase.GetOneQuiz(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getOneQuiz.FromDomain(clean))
}

func (ctr *QuizsController) CalculateScoreQuiz(c echo.Context) error {
	var req requests.RequestJawabans

	token, err := middlewares.ExtractClaims(c)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}
	ctx := c.Request().Context()
	req.ModulId = c.Param("modulid")

	clean, err := ctr.QuizsUsecase.CalculateScoreQuiz(ctx, requests.ToListDomain(req), token.Userid)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, jawabanQuiz.FromDomain(clean))
}
