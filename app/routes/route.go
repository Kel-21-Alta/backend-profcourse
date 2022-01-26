package routes

import (
	"profcourse/controllers/courses"
	"profcourse/controllers/feedback"
	"profcourse/controllers/materies"
	"profcourse/controllers/moduls"
	"profcourse/controllers/quizs"
	requestusers "profcourse/controllers/request_users"
	"profcourse/controllers/spesializations"
	"profcourse/controllers/summary"
	"profcourse/controllers/users"
	"profcourse/controllers/users_courses"
	"profcourse/controllers/users_spesializations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware                  middleware.JWTConfig
	UserController                 users.UserController
	CourseController               courses.CourseController
	UserCourseController           users_courses.UsersCoursesController
	ModulController                moduls.ModulController
	SummaryController              summary.SummaryController
	SpesializationController       spesializations.SpesializationController
	MateriesController             materies.MateriesController
	QuizController                 quizs.QuizsController
	FeedbackController             feedback.FeedbackController
	RequestUserController          requestusers.RequestUserController
	UsersSpesializationsController users_spesializations.UsersSpesializationController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	ev1 := e.Group("api/v1/")
	ev1.POST("login", cl.UserController.Login)
	ev1.POST("admin/login", cl.UserController.LoginAdmin)
	ev1.PUT("forgetpassword", cl.UserController.ForgetPassword)

	withJWT := ev1.Group("")
	withJWT.Use(middleware.JWTWithConfig(cl.JWTMiddleware))

	// User
	withJWT.POST("users", cl.UserController.CreateUser)
	withJWT.DELETE("users/:userid", cl.UserController.DeleteUser)
	withJWT.PUT("users/:userid", cl.UserController.UpdateUser)
	withJWT.GET("currentuser", cl.UserController.GetCurrentUser)
	withJWT.GET("users/:userid", cl.UserController.GetDetailUser)
	withJWT.PUT("currentuser", cl.UserController.UpdateCurrentUserFromUser)
	withJWT.PUT("changepassword", cl.UserController.ChangePassword)
	withJWT.GET("users", cl.UserController.GetAllUser)
	withJWT.GET("users/reports/:userid", cl.UserController.GenerateReportUser)

	// course
	withJWT.POST("courses", cl.CourseController.CreateCourse)
	withJWT.GET("courses/:courseid", cl.CourseController.GetOneCourse)
	withJWT.PUT("courses/:courseid", cl.CourseController.UpdateCourse)
	withJWT.DELETE("courses/:courseid", cl.CourseController.DeleteCourse)
	withJWT.GET("courses", cl.CourseController.GetAllCourses)
	withJWT.GET("coursesendroll", cl.UserCourseController.GetUserCourseEndroll)
	withJWT.GET("user/courses", cl.CourseController.GetAllCoursesUser)

	// untuk melakukan register course user
	withJWT.POST("course/register", cl.UserCourseController.UserRegisterCourse)

	// update progress
	withJWT.PUT("materi/progress", cl.MateriesController.UpdateProgressMateri)

	//Modul
	withJWT.POST("moduls", cl.ModulController.CreateModul)
	withJWT.GET("moduls/course/:courseid", cl.ModulController.GetAllModulCourse)
	withJWT.PUT("moduls/:modulid", cl.ModulController.UpdateModul)
	withJWT.DELETE("moduls/:modulid", cl.ModulController.DeleteModul)

	//Materi
	withJWT.POST("materi", cl.MateriesController.CreateMateries)
	withJWT.GET("moduls/:modulid", cl.MateriesController.GetAllMateri)
	withJWT.DELETE("materi/:materiid", cl.MateriesController.DeleteMateries)
	withJWT.PUT("materi/:materiid", cl.MateriesController.UpdateMateri)
	withJWT.GET("materi/:materiid", cl.MateriesController.GetOneMateri)

	//Quiz
	withJWT.POST("quizs", cl.QuizController.CreateQuiz)
	withJWT.PUT("quizs/:quizid", cl.QuizController.UpdateQuiz)
	withJWT.DELETE("quizs/:quizid", cl.QuizController.DeleteQuiz)
	withJWT.GET("quizs/modul/:modulid", cl.QuizController.GetAllQuizModul)
	withJWT.GET("quizs/:quizid", cl.QuizController.GetOneQuiz)
	withJWT.POST("quizs/modul/:modulid", cl.QuizController.CalculateScoreQuiz)

	withJWT.GET("summary", cl.SummaryController.GetAllSummary)

	// spesailization
	withJWT.POST("spesializations", cl.SpesializationController.CreateSpesialization)
	withJWT.GET("spesializations/:spesializationid", cl.SpesializationController.GetOneSpesialization)
	withJWT.GET("spesializations", cl.SpesializationController.GetAllSpesialization)

	//user spesialization
	withJWT.POST("spesialization/register", cl.UsersSpesializationsController.RegisterSpesialization)

	//	feedback
	withJWT.POST("feedback", cl.FeedbackController.CreateFeedback)
	withJWT.GET("feedback/course/:courseid", cl.FeedbackController.GetAllFeedbackByCourse)
	withJWT.DELETE("feedback/:feedbackid", cl.FeedbackController.DeleteFeedback)

	// Request User
	withJWT.POST("requestusers", cl.RequestUserController.CreateRequest)
	withJWT.GET("requestusers", cl.RequestUserController.GetAllRequestUser)
	withJWT.GET("admin/requestusers", cl.RequestUserController.AdminGetAllRequestUser)
	withJWT.DELETE("requestusers/:requestusers", cl.RequestUserController.DeleteRequestUser)
	withJWT.PUT("requestusers/:requestusers", cl.RequestUserController.UpdateRequestUser)
	withJWT.GET("requestusers/:requestusers", cl.RequestUserController.GetOneRequestUser)
	withJWT.GET("categoryrequestuser", cl.RequestUserController.GetAllCategoryRequest)
}
