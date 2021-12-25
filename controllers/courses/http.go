package courses

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/business/courses"
	controller "profcourse/controllers"
	"profcourse/controllers/courses/requests"
	"profcourse/controllers/courses/responses/createCourse"
)

type CourseController struct {
	CourseUsecase courses.Usecase
}

func NewCourseController(cc courses.Usecase) *CourseController {
	return &CourseController{CourseUsecase: cc}
}

func (cc CourseController) CreateCourse(c echo.Context) error {
	ctx := c.Request().Context()
	var err error

	req := requests.CreateCourseRequest{}
	if err = c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	req.FileImage, err = c.FormFile("file_image")
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	clean, err := cc.CourseUsecase.CreateCourse(ctx, req.ToDomain())
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusCreated, createCourse.FromDomain(clean))
}
