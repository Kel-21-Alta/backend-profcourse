package routes

import (
	"github.com/labstack/echo/v4"
	"profcourse/controllers/users"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	ev1 := e.Group("api/v1/")
	ev1.POST("users", cl.UserController.CreateUser)
	ev1.POST("login", cl.UserController.Login)
	ev1.PUT("forgetpassword", cl.UserController.ForgetPassword)
}
