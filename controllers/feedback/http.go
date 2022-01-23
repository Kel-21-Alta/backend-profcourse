package feedback

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/feedback"
	controller "profcourse/controllers"
	"profcourse/controllers/feedback/requests"
	"profcourse/controllers/feedback/responses/createFeedback"
	"profcourse/controllers/feedback/responses/getAllFeedbackCourse"
)

type FeedbackController struct {
	FeedbackUsecase feedback.Usecase
}

func NewFeedbackController(feedback feedback.Usecase) *FeedbackController {
	return &FeedbackController{FeedbackUsecase: feedback}
}

func (ctr *FeedbackController) CreateFeedback(c echo.Context) error {
	var req requests.CreateFeedbackRequest

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}
	domain := req.Todomain()

	domain.UserId = token.Userid

	ctx := c.Request().Context()
	result, err := ctr.FeedbackUsecase.CreateFeedback(ctx, domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusCreated, createFeedback.FromDomain(result))
}

func (ctr *FeedbackController) GetAllFeedbackByCourse(c echo.Context) error {
	var domain feedback.CourseReviews

	domain.CourseId = c.Param("courseid")

	ctx := c.Request().Context()
	result, err := ctr.FeedbackUsecase.GetAllFeedbackByCourse(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	response := getAllFeedbackCourse.FromListDomain(result)

	return controller.NewResponseSuccess(c, http.StatusOK, response)
}
