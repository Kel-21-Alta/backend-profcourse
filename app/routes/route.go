package routes

import (
	"profcourse/controllers/courses"
	"profcourse/controllers/moduls"
	"profcourse/controllers/spesializations"
	"profcourse/controllers/summary"
	"profcourse/controllers/users"
	"profcourse/controllers/users_courses"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	withJWT.PUT("users/:userid", cl.UserController.UpdateUser)
	withJWT.GET("currentuser", cl.UserController.GetCurrentUser)
	withJWT.PUT("changepassword", cl.UserController.ChangePassword)

	// course
	withJWT.POST("courses", cl.CourseController.CreateCourse)
	withJWT.GET("courses/:courseid", cl.CourseController.GetOneCourse)
	withJWT.PUT("courses/:courseid", cl.CourseController.UpdateCourse)
	withJWT.DELETE("courses/:courseid", cl.CourseController.DeleteCourse)
	withJWT.GET("courses", cl.CourseController.GetAllCourses)

	// untuk melakukan register course user
	withJWT.POST("course/register", cl.UserCourseController.UserRegisterCourse)

	//Modul
	withJWT.POST("moduls", cl.ModulController.CreateModul)
	withJWT.GET("moduls/:modulid", cl.ModulController.GetOneModul)
	withJWT.PUT("moduls/:modulid", cl.ModulController.UpdateModul)
	withJWT.DELETE("moduls/:modulid", cl.ModulController.DeleteModul)

	withJWT.GET("summary", cl.SummaryController.GetAllSummary)

	// spesailization
	withJWT.POST("spesializations", cl.SpesializationController.CreateSpesialization)
	withJWT.GET("spesializations/:spesializationid", cl.SpesializationController.GetOneSpesialization)
	withJWT.GET("spesializations", cl.SpesializationController.GetAllSpesialization)

}
