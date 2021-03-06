package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"profcourse/app/middlewares"
	"profcourse/app/routes"
	_coursesUsecase "profcourse/business/courses"
	_feedbackUsecase "profcourse/business/feedback"
	_materiesUsecase "profcourse/business/materies"
	_modulsUsecase "profcourse/business/moduls"
	_quizsUsecase "profcourse/business/quizs"
	_requestUsersUsecase "profcourse/business/request_users"
	_spesializationUsecase "profcourse/business/spesializations"
	_summaryUsecase "profcourse/business/summary"
	_userUsecase "profcourse/business/users"
	_usersCourseUsercase "profcourse/business/users_courses"
	_userSpesializationsUsecase "profcourse/business/users_spesializations"
	"profcourse/controllers/courses"
	"profcourse/controllers/feedback"
	"profcourse/controllers/materies"
	_modulController "profcourse/controllers/moduls"
	"profcourse/controllers/quizs"
	requestusers "profcourse/controllers/request_users"
	_spesializationsController "profcourse/controllers/spesializations"
	_summaryController "profcourse/controllers/summary"
	_userController "profcourse/controllers/users"
	_usersCourseController "profcourse/controllers/users_courses"
	"profcourse/controllers/users_spesializations"
	_driversFectory "profcourse/drivers"
	_coursesMysqlRepo "profcourse/drivers/databases/courses"
	_feedbackMysqlRepo "profcourse/drivers/databases/feedback"
	_materiesMysqlRepo "profcourse/drivers/databases/materies"
	_modulsMysqlRepo "profcourse/drivers/databases/moduls"
	_quizsMysqlRepo "profcourse/drivers/databases/quizs"
	_requestuserMysqlRepo "profcourse/drivers/databases/request_users"
	_spesializationMysqlRepo "profcourse/drivers/databases/spesialization"
	_userMysqlRepo "profcourse/drivers/databases/users"
	_usersCourseMysqlRepo "profcourse/drivers/databases/users_courses"
	_usersSpesializationMysqlRepo "profcourse/drivers/databases/users_spesializations"
	_dbDriver "profcourse/drivers/mysql"
	"profcourse/drivers/thirdparties/mail"
	"time"
)

func init() {
	viper.SetConfigFile("./config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool("debug") {
		log.Println("Server RUN on DEBUG mode")
	}
}

func DbMigration(db *gorm.DB) {
	var err error
	err = db.AutoMigrate(
		&_userMysqlRepo.User{},
		&_coursesMysqlRepo.Courses{},
		&_usersCourseMysqlRepo.UsersCourses{},
		&_modulsMysqlRepo.Moduls{},
		&_spesializationMysqlRepo.Spesialization{},
		&_materiesMysqlRepo.Materi{},
		&_materiesMysqlRepo.MateriUserComplate{},
		&_quizsMysqlRepo.Quiz{},
		&_quizsMysqlRepo.PilihanQuiz{},
		&_modulsMysqlRepo.SkorUserModul{},
		&_feedbackMysqlRepo.Feedback{},
		&_requestuserMysqlRepo.RequestUser{},
		&_requestuserMysqlRepo.CategoryRequest{},
		&_usersSpesializationMysqlRepo.UsersSpesializations{},
	)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Mingration Success")
	}
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}

	congfigSmtp := mail.Email{
		Host:       viper.GetString("smtp.host"),
		Port:       viper.GetInt("smtp.port"),
		SenderName: viper.GetString("smtp.name"),
		AuthEmail:  viper.GetString("smtp.email"),
		Password:   viper.GetString("smtp.password"),
	}

	configJwt := middlewares.ConfigJwt{
		SecretJwt:       viper.GetString("jwt.secret"),
		ExpiredDuration: viper.GetInt("jwt.expired"),
	}

	conn := configDB.InitialDB()
	DbMigration(conn)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Static("/public", "public")
	e.Static("/pdf", "")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	timeout := time.Duration(viper.GetInt("context.timeout")) * time.Second

	smtpRepository := _driversFectory.NewSmtpRepository(congfigSmtp)
	pdfGenerater := _driversFectory.NewGeneratePDFRepostUser()

	mysqlUserRepository := _driversFectory.NewMysqlUserRepository(conn)
	userUsecase := _userUsecase.NewUserUsecase(mysqlUserRepository, timeout, smtpRepository, pdfGenerater, configJwt)
	userCtrl := _userController.NewUserController(userUsecase)

	mysqlCourseRepository := _driversFectory.NewMysqlCourseRepository(conn)
	courseUsecase := _coursesUsecase.NewCourseUseCase(mysqlCourseRepository, timeout)
	couserCtrl := courses.NewCourseController(courseUsecase)

	mysqlUserCourseRepository := _driversFectory.NewMysqlUserCourseRepository(conn)
	userCourseUsecase := _usersCourseUsercase.NewUsersCoursesUsecase(mysqlUserCourseRepository, userUsecase, courseUsecase, timeout)
	userCourseController := _usersCourseController.NewUsesrCoursesController(userCourseUsecase)

	mysqlSpesializationRepository := _driversFectory.NewMysqlSpesializationRepository(conn)
	spesializationUsecae := _spesializationUsecase.NewSpesializationUsecase(mysqlSpesializationRepository, timeout)
	spesializationController := _spesializationsController.NewSpesializationController(spesializationUsecae)

	summaryUsecase := _summaryUsecase.NewSummaryUsecase(timeout, courseUsecase, userUsecase, spesializationUsecae)
	summaryController := _summaryController.NewSummaryController(summaryUsecase)

	mysqlModulRepository := _driversFectory.NewMysqlModulRepository(conn)
	modulUsecase := _modulsUsecase.NewModulUsecase(mysqlModulRepository, courseUsecase, timeout)
	modulCtrl := _modulController.NewModulsController(modulUsecase)

	mysqlMateriesRepository := _driversFectory.NewMysqlMateriesRepository(conn)
	materiesUsecase := _materiesUsecase.NewMateriesUsecase(mysqlMateriesRepository, userCourseUsecase, modulUsecase, timeout)
	materiesController := materies.NewMateriesController(materiesUsecase)

	myzqlQuizRepository := _driversFectory.NewMysqlQuizsRepository(conn)
	quizsUsecase := _quizsUsecase.NewQuizUsecase(myzqlQuizRepository, modulUsecase, userCourseUsecase, timeout)
	quizController := quizs.NewQuizsController(quizsUsecase)

	mysqlFeedbackRepository := _driversFectory.NewMysqlFeedbackRepository(conn)
	feedbackUsecase := _feedbackUsecase.NewFeedbackUsecase(mysqlFeedbackRepository, timeout)
	feedbackController := feedback.NewFeedbackController(feedbackUsecase)

	mysqlRequestUserRepo := _driversFectory.NewRequestUserRepository(conn)
	requestUsecase := _requestUsersUsecase.NewRequestUserUsecase(mysqlRequestUserRepo, timeout)
	requestUserController := requestusers.NewRequestUserController(requestUsecase)

	mysqlUserSpesializationRepo := _driversFectory.NewUsersSpesializationsRepository(conn)
	usersSpesializationUsecase := _userSpesializationsUsecase.NewUsersSpesializationsUsecase(mysqlUserSpesializationRepo, spesializationUsecae, userCourseUsecase, timeout)
	usersSpesializationController := users_spesializations.NewUsersSpesializationController(usersSpesializationUsecase)

	routesInit := routes.ControllerList{
		UserController:           *userCtrl,
		CourseController:         *couserCtrl,
		JWTMiddleware:            configJwt.Init(),
		UserCourseController:     *userCourseController,
		ModulController:          *modulCtrl,
		SummaryController:        *summaryController,
		SpesializationController: *spesializationController,
		MateriesController:       *materiesController,
		QuizController:           *quizController,
		FeedbackController:       *feedbackController,
		RequestUserController:    *requestUserController,
		UsersSpesializationsController: *usersSpesializationController,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
