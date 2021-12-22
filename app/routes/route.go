package routes

import (
	"github.com/labstack/echo/v4"
	"profcourse/controllers/users"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("/api/v1/users", cl.UserController.CreateUser)
}
