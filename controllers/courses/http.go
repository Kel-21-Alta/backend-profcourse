package courses

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/courses"
	controller "profcourse/controllers"
	"profcourse/controllers/courses/requests"
	"profcourse/controllers/courses/responses/createCourse"
	"profcourse/controllers/courses/responses/deletecourse"
	"profcourse/controllers/courses/responses/getAllCourses"
	"profcourse/controllers/courses/responses/getOneCourse"
	"profcourse/controllers/courses/responses/updateCourse"
	"strconv"
)

type CourseController struct {
	CourseUsecase courses.Usecase
}

func NewCourseController(cc courses.Usecase) *CourseController {
	return &CourseController{CourseUsecase: cc}
}

func (cc CourseController) GetOneCourse(c echo.Context) error {
	ctx := c.Request().Context()
	var domain courses.Domain
	token, _ := middlewares.ExtractClaims(c)

	domain.ID = c.Param("courseid")
	domain.InfoUser.CurrentUser = token.Userid

	clean, err := cc.CourseUsecase.GetOneCourse(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getOneCourse.FromDomain(clean))
}

func (cc CourseController) CreateCourse(c echo.Context) error {
	ctx := c.Request().Context()
	token, _ := middlewares.ExtractClaims(c)
	var err error

	req := requests.CreateCourseRequest{}
	if err = c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	req.UserId = token.Userid
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	clean, err := cc.CourseUsecase.CreateCourse(ctx, req.ToDomain())
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusCreated, createCourse.FromDomain(clean))
}

func (cc CourseController) GetAllCourses(c echo.Context) error {
	ctx := c.Request().Context()
	var domain courses.Domain
	var err error

	domain.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	domain.Offset, _ = strconv.Atoi(c.QueryParam("offset"))
	domain.Sort = c.QueryParam("sort")
	domain.SortBy = c.QueryParam("sortby")
	domain.KeywordSearch = c.QueryParam("s")
	domain.ParamStatus, _ = strconv.Atoi(c.QueryParam("status"))

	// Usecase
	clean, err := cc.CourseUsecase.GetAllCourses(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getAllCourses.FromListDomain(clean))
}

func (cc CourseController) UpdateCourse(c echo.Context) error {

	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	req := requests.UpdateCourse{}
	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}
	ctx := c.Request().Context()
	domain := req.ToDomain()
	domain.ID = c.Param("courseid")
	clean, err := cc.CourseUsecase.UpdateCourse(ctx, domain, &courses.Token{UserId: token.Userid, Role: token.Role})

	if err != nil {
		return controller.NewResponseError(c, err)
	}
	return controller.NewResponseSuccess(c, http.StatusOK, updateCourse.FromDomain(clean))
}

func (cc CourseController) DeleteCourse(c echo.Context) error {
	var id string
	var err error
	var token *middlewares.JwtCustomClaims
	token, err = middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}
	id = c.Param("courseid")
	ctx := c.Request().Context()
	_, err = cc.CourseUsecase.DeleteCourse(ctx, id, courses.Token{
		UserId: token.Userid,
		Role:   token.Role,
	})

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, deletecourse.DeleteCourseResponse{Message: "Kursus berhasil dihapus"})
}

func (cc CourseController) GetAllCoursesUser(c echo.Context) error {
	var domain courses.Domain

	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	domain.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	domain.Offset, _ = strconv.Atoi(c.QueryParam("offset"))
	domain.Sort = c.QueryParam("sort")
	domain.KeywordSearch = c.QueryParam("s")
	domain.ParamStatus, _ = strconv.Atoi(c.QueryParam("status"))
	domain.TeacherId = token.Userid

	ctx := c.Request().Context()
	result, err := cc.CourseUsecase.GetAllCourseUser(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getAllCourses.FromListDomain(&result))
}
