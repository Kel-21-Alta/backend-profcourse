package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"profcourse/app/middlewares"
	"profcourse/business/users"
	controller "profcourse/controllers"
	"profcourse/controllers/users/requests"
	"profcourse/controllers/users/reseponses/changePassword"
	"profcourse/controllers/users/reseponses/currentUser"
	"profcourse/controllers/users/reseponses/deleteUser"
	"profcourse/controllers/users/reseponses/forgetPassword"
	"profcourse/controllers/users/reseponses/login"
	"profcourse/controllers/users/reseponses/userCreated"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{userUsecase: uc}
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
