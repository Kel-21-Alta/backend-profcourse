package users

import (
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/users"
	controller "profcourse/controllers"
	"profcourse/controllers/users/requests"
	"profcourse/controllers/users/reseponses/changePassword"
	"profcourse/controllers/users/reseponses/currentUser"
	"profcourse/controllers/users/reseponses/deleteUser"
	"profcourse/controllers/users/reseponses/forgetPassword"
	"profcourse/controllers/users/reseponses/getAllUser"
	"profcourse/controllers/users/reseponses/login"
	"profcourse/controllers/users/reseponses/updateUser"
	"profcourse/controllers/users/reseponses/userCreated"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{userUsecase: uc}
}

func (ctrl UserController) GetAllUser(c echo.Context) error {
	var domain users.Domain

	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	domain.Role = token.Role
	domain.Query.Search = c.QueryParam("s")
	domain.Query.Sort = c.QueryParam("sort")
	domain.Query.SortBy = c.QueryParam("sortby")
	domain.Query.Offset, _ = strconv.Atoi(c.QueryParam("offset"))
	domain.Query.Limit, _ = strconv.Atoi(c.QueryParam("limit"))

	ctx := c.Request().Context()
	result, err := ctrl.userUsecase.GetAllUser(ctx, &domain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, getAllUser.FromListDomain(result))
}

func (ctrl UserController) DeleteUser(c echo.Context) error {
	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	domain := users.Domain{
		ID:     token.Userid,
		Role:   token.Role,
		IdUser: c.Param("userid"),
	}

	ctx := c.Request().Context()
	user, err := ctrl.userUsecase.DeleteUser(ctx, domain)
	if err != nil {
		return err
	}
	return controller.NewResponseSuccess(c, http.StatusOK, deleteUser.FromDomain(user))
}

func (ctrl UserController) UpdateUser(c echo.Context) error {
	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	ctx := c.Request().Context()
	req := requests.UpdateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}
	domain := req.ToDomain()
	domain.ID = token.Userid
	domain.Role = token.Role

	user, err := ctrl.userUsecase.UpdateUser(ctx, *domain)
	if err != nil {
		return err
	}
	return controller.NewResponseSuccess(c, http.StatusOK, updateUser.FromDomain(user))
}

func (ctrl *UserController) CreateUser(c echo.Context) error {

	tokenJwt, _ := middlewares.ExtractClaims(c)

	if tokenJwt.Role != int8(1) {
		return controller.NewResponseError(c, controller.FORBIDDIN_USER)
	}

	ctx := c.Request().Context()
	req := requests.UserRequest{}

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	clean, err := ctrl.userUsecase.CreateUser(ctx, *req.ToDomain())

	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusCreated, userCreated.FromDomain(clean))
}

func (ctrl *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := requests.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	clean, err := ctrl.userUsecase.Login(ctx, *req.ToDomain())
	if err != nil {
		return controller.NewResponseError(c, err)
	}
	return controller.NewResponseSuccess(c, http.StatusOK, login.FromDomain(clean))
}

func (ctrl *UserController) LoginAdmin(c echo.Context) error {
	ctx := c.Request().Context()
	req := requests.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	clean, err := ctrl.userUsecase.LoginAdmin(ctx, *req.ToDomain())
	if err != nil {
		return controller.NewResponseError(c, err)
	}
	return controller.NewResponseSuccess(c, http.StatusOK, login.FromDomain(clean))
}

func (ctrl *UserController) ForgetPassword(c echo.Context) error {
	ctx := c.Request().Context()
	req := requests.ForgetPasswordRequest{}
	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	_, err := ctrl.userUsecase.ForgetPassword(ctx, *req.ToDomain())
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, forgetPassword.GenerateResponses())
}

func (ctrl *UserController) GetCurrentUser(c echo.Context) error {
	tokenJwt, _ := middlewares.ExtractClaims(c)

	ctx := c.Request().Context()
	userDomain := users.Domain{ID: tokenJwt.Userid}
	clean, err := ctrl.userUsecase.GetCurrentUser(ctx, userDomain)

	if err != nil {
		return controller.NewResponseError(c, err)
	}
	return controller.NewResponseSuccess(c, http.StatusOK, currentUser.FromDomain(clean))
}

func (ctrl *UserController) ChangePassword(c echo.Context) error {
	tokenJwt, _ := middlewares.ExtractClaims(c)
	req := requests.ChangePasswordRequest{}
	if err := c.Bind(&req); err != nil {

		return controller.NewResponseError(c, err)
	}

	ctx := c.Request().Context()

	domain := req.ToDomain()
	domain.ID = tokenJwt.Userid
	_, err := ctrl.userUsecase.ChangePassword(ctx, domain)
	if err != nil {
		return err
	}

	return controller.NewResponseSuccess(c, http.StatusOK, changePassword.GenerateMessage())
}

func (ctrl *UserController) UpdateCurrentUserFromUser(c echo.Context) error {
	ctx := c.Request().Context()
	token, err := middlewares.ExtractClaims(c)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	var req requests.UpdateCurrentUser

	if err := c.Bind(&req); err != nil {
		return controller.NewResponseError(c, err)
	}

	var domain *users.Domain = req.ToDomain()

	domain.ID = token.Userid
	result, err := ctrl.userUsecase.UpdateDataCurrentUser(ctx, domain)
	if err != nil {
		return controller.NewResponseError(c, err)
	}

	return controller.NewResponseSuccess(c, http.StatusOK, updateUser.FromDomain(result))
}
