package users_courses

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/users_courses"
	controller "profcourse/controllers"
	"profcourse/controllers/users_courses/request"
	"profcourse/controllers/users_courses/responses/courseUserEndroll"
	"profcourse/controllers/users_courses/responses/userRegisterCourse"
)

type UsersCoursesController struct {
	UsersCoursesUsecase users_courses.Usecase
}

func NewUsesrCoursesController(ucc users_courses.Usecase) *UsersCoursesController {
	return &UsersCoursesController{UsersCoursesUsecase: ucc}
}

func (uc *UsersCoursesController) UserRegisterCourse(c echo.Context) error {
	var err error
	var token *middlewares.JwtCustomClaims

	ctx := c.Request().Context()
	req := request.RegisterCourseRequest{}
	token, err = middlewares.ExtractClaims(c)

	if err != nil {
		return controller.NewResponseError(c, err)
	}
	if err = c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	domain := req.ToDomain()

	domain.UserId = token.Userid
	_, err = uc.UsersCoursesUsecase.UserRegisterCourse(ctx, domain)
	if err != nil {
		return controller.NewResponseError(c, err)
	}
	return controller.NewResponseSuccess(c, http.StatusOK, userRegisterCourse.UserRegisterCourseResponse{Message: "Berhasil mendaftarkan user pada kursus"})
}

func (uc *UsersCoursesController) GetUserCourseEndroll(c echo.Context) error {
	var req users_courses.User

	token, err := middlewares.ExtractClaims(c)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	req.UserID = token.Userid
	ctx := c.Request().Context()
	result, err := uc.UsersCoursesUsecase.GetUserCourseEndroll(ctx, &req)
	if err != nil {
		return controller.NewResponseError(c, err)
	}
	return controller.NewResponseSuccess(c, http.StatusOK, courseUserEndroll.FromDomain(result))
}
