package requestusers

import (
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/request_users"
	controller "profcourse/controllers"
	"profcourse/controllers/request_users/requests"
	createrequestuser "profcourse/controllers/request_users/responses/createRequestuser"
	getAllCategoryRequestUser "profcourse/controllers/request_users/responses/getAllCategoryRequestUser"

	"github.com/labstack/echo/v4"
)

type RequestUserController struct {
	RequestUserUsecase request_users.Usecase
}

func NewRequestUserController(usecase request_users.Usecase) *RequestUserController {
	return &RequestUserController{RequestUserUsecase: usecase}
}

func (ctr *RequestUserController) CreateRequest(c echo.Context) error {

	var req requests.CreateRequestUser

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	token, err := middlewares.ExtractClaims(c)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	var domain = req.ToDomain()
	domain.UserId = token.Userid

	ctx := c.Request().Context()
	result, err := ctr.RequestUserUsecase.CreateRequest(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, createrequestuser.FromDomain(result))
}

func (ctr *RequestUserController) GetAllCategoryRequest(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := ctr.RequestUserUsecase.GetAllCategoryRequest(ctx)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getAllCategoryRequestUser.FromListDomain(result))
}

func (ctr *RequestUserController) GetAllRequestCategory(c echo.Context) error {

	ctx := c.Request().Context()

	var domain request_users.Domain

	domain.Query.Sort = c.QueryParam("sort")
	domain.Query.Offset = c.QueryParam("offset")
	domain.Query.Limit = c.QueryParam("limit")
	domain.Query.Search = c.QueryParam("s")

	result, err := ctr.RequestUserUsecase.GetAllRequestUser(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, result)
}