package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"profcourse/controllers/courses"
	"profcourse/controllers/moduls"
	"profcourse/controllers/users"
	"profcourse/controllers/users_courses"
)

type ControllerList struct {
	JWTMiddleware        middleware.JWTConfig
	UserController       users.UserController
	CourseController     courses.CourseController
	UserCourseController users_courses.UsersCoursesController
	ModulController      moduls.ModulController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	configCors := middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}

	ev1 := e.Group("api/v1/")
	ev1.Use(middleware.CORSWithConfig(configCors))
	ev1.POST("login", cl.UserController.Login)
	ev1.PUT("forget-password", cl.UserController.ForgetPassword)

	withJWT := ev1.Group("")
	withJWT.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	withJWT.Use(middleware.CORSWithConfig(configCors))
	withJWT.POST("users", cl.UserController.CreateUser)
	withJWT.GET("currentuser", cl.UserController.GetCurrentUser)
	withJWT.PUT("changepassword", cl.UserController.ChangePassword)

	withJWT.POST("courses", cl.CourseController.CreateCourse)
	withJWT.GET("courses/:courseid", cl.CourseController.GetOneCourse)
	withJWT.GET("courses", cl.CourseController.GetAllCourses)

	withJWT.POST("course/register", cl.UserCourseController.UserRegisterCourse)

	withJWT.POST("modul", cl.ModulController.CreateModul)
}
