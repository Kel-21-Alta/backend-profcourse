package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"profcourse/controllers/courses"
	"profcourse/controllers/users"
)

type ControllerList struct {
	JWTMiddleware    middleware.JWTConfig
	UserController   users.UserController
	CourseController courses.CourseController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	ev1 := e.Group("api/v1/")
	ev1.POST("login", cl.UserController.Login)
	ev1.PUT("forget-password", cl.UserController.ForgetPassword)

	withJWT := ev1.Group("")
	withJWT.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	withJWT.POST("users", cl.UserController.CreateUser)
	withJWT.GET("currentuser", cl.UserController.GetCurrentUser)
	withJWT.PUT("changepassword", cl.UserController.ChangePassword)

	withJWT.POST("courses", cl.CourseController.CreateCourse)
	withJWT.GET("courses/:courseid", cl.CourseController.GetOneCourse)
}
