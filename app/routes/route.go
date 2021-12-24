package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"profcourse/controllers/users"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	ev1 := e.Group("api/v1/")
	ev1.POST("login", cl.UserController.Login)
	ev1.PUT("forget-password", cl.UserController.ForgetPassword)

	withJWT := ev1.Group("")
	withJWT.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	withJWT.POST("users", cl.UserController.CreateUser)
	withJWT.GET("currentuser", cl.UserController.GetCurrentUser)

}
