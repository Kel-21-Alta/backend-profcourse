package requestusers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/request_users"
	controller "profcourse/controllers"
	"profcourse/controllers/request_users/requests"
	createrequestuser "profcourse/controllers/request_users/responses/createRequestuser"
	getAllCategoryRequestUser "profcourse/controllers/request_users/responses/getAllCategoryRequestUser"
	"profcourse/controllers/request_users/responses/getAllRequestUser"
	"profcourse/controllers/request_users/responses/getOneRequestUser"
	"profcourse/controllers/request_users/responses/updateRequest"
	"strconv"
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

func (ctr *RequestUserController) GetAllRequestUser(c echo.Context) error {

	ctx := c.Request().Context()

	var domain request_users.Domain

	token, err := middlewares.ExtractClaims(c)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	domain.UserId = token.Userid
	domain.Query.Sort = c.QueryParam("sort")
	domain.Query.Offset, _ = strconv.Atoi(c.QueryParam("offset"))
	domain.Query.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	domain.Query.Search = c.QueryParam("s")

	result, err := ctr.RequestUserUsecase.GetAllRequestUser(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getAllRequestUser.FromListDomain(result))
}

func (ctr *RequestUserController) DeleteRequestUser(c echo.Context) error {

	var domain request_users.Domain

	domain.Id = c.Param("requestusers")

	ctx := c.Request().Context()
	_, err := ctr.RequestUserUsecase.DeleteRequestUser(ctx, &domain)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	type Message struct {
		 Message string
	}
	return controller.NewResponseSuccess(c, http.StatusOK, Message{Message: "Request dengan id :" + domain.Id + " telah terhapus"} )
}

func (ctr *RequestUserController) UpdateRequestUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req requests.UpdateRequestUser

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	var domain = req.ToDomain()

	token, err := middlewares.ExtractClaims(c)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	domain.Id = c.Param("requestusers")
	domain.UserId = token.Userid

	result, err := ctr.RequestUserUsecase.UpdateRequestUser(ctx, &domain)
	if err != nil {
		return controller.NewResponseError(c, err)

	}

	return controller.NewResponseSuccess(c, http.StatusOK, updateRequest.FromDomain(result))
}

func (ctr *RequestUserController) GetOneRequestUser(c echo.Context) error {
	ctx := c.Request().Context()

	var domain = request_users.Domain{Id: c.Param("requestusers")}

	result, err := ctr.RequestUserUsecase.GetOneRequestUser(ctx, &domain)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getOneRequestUser.FromDomain(result))
}