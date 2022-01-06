package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"profcourse/controllers/courses"
	"profcourse/controllers/moduls"
	"profcourse/controllers/spesializations"
	"profcourse/controllers/summary"
	"profcourse/controllers/users"
	"profcourse/controllers/users_courses"
)

type ControllerList struct {
	JWTMiddleware            middleware.JWTConfig
	UserController           users.UserController
	CourseController         courses.CourseController
	UserCourseController     users_courses.UsersCoursesController
	ModulController          moduls.ModulController
	SummaryController        summary.SummaryController
	SpesializationController spesializations.SpesializationController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	ev1 := e.Group("api/v1/")
	ev1.POST("login", cl.UserController.Login)
	ev1.POST("admin/login", cl.UserController.LoginAdmin)
	ev1.PUT("forgetpassword", cl.UserController.ForgetPassword)

	withJWT := ev1.Group("")
	withJWT.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	withJWT.POST("users", cl.UserController.CreateUser)
	withJWT.DELETE("users/:userid", cl.UserController.DeleteUser)
	withJWT.GET("currentuser", cl.UserController.GetCurrentUser)
	withJWT.PUT("changepassword", cl.UserController.ChangePassword)

	withJWT.POST("courses", cl.CourseController.CreateCourse)
	withJWT.GET("courses/:courseid", cl.CourseController.GetOneCourse)
	withJWT.GET("courses", cl.CourseController.GetAllCourses)

	withJWT.POST("course/register", cl.UserCourseController.UserRegisterCourse)

	withJWT.POST("modul", cl.ModulController.CreateModul)

	withJWT.GET("summary", cl.SummaryController.GetAllSummary)

	withJWT.POST("spesializations", cl.SpesializationController.CreateSpesialization)
	withJWT.GET("spesializations", cl.SpesializationController.GetAllSpesialization)
}
