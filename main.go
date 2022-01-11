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
	_modulsUsecase "profcourse/business/moduls"
	_spesializationUsecase "profcourse/business/spesializations"
	_summaryUsecase "profcourse/business/summary"
	_userUsecase "profcourse/business/users"
	_usersCourseUsercase "profcourse/business/users_courses"
	"profcourse/controllers/courses"
	_modulController "profcourse/controllers/moduls"
	_spesializationsController "profcourse/controllers/spesializations"
	_summaryController "profcourse/controllers/summary"
	_userController "profcourse/controllers/users"
	_usersCourseController "profcourse/controllers/users_courses"
	_driversFectory "profcourse/drivers"
	_coursesMysqlRepo "profcourse/drivers/databases/courses"
	_materiesMysqlRepo "profcourse/drivers/databases/materies"
	_modulsMysqlRepo "profcourse/drivers/databases/moduls"
	_spesializationMysqlRepo "profcourse/drivers/databases/spesialization"
	_userMysqlRepo "profcourse/drivers/databases/users"
	_usersCourseMysqlRepo "profcourse/drivers/databases/users_courses"
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

	timeout := time.Duration(viper.GetInt("context.timeout")) * time.Second

	smtpRepository := _driversFectory.NewSmtpRepository(congfigSmtp)

	mysqlUserRepository := _driversFectory.NewMysqlUserRepository(conn)
	userUsecase := _userUsecase.NewUserUsecase(mysqlUserRepository, timeout, smtpRepository, configJwt)
	userCtrl := _userController.NewUserController(userUsecase)

	mysqlCourseRepository := _driversFectory.NewMysqlCourseRepository(conn)
	courseUsecase := _coursesUsecase.NewCourseUseCase(mysqlCourseRepository, timeout)
	couserCtrl := courses.NewCourseController(courseUsecase)

	mysqlUserCourseRepository := _driversFectory.NewMysqlUserCourseRepository(conn)
	userCourseUsecase := _usersCourseUsercase.NewUsersCoursesUsecase(mysqlUserCourseRepository, timeout)
	userCourseController := _usersCourseController.NewUsesrCoursesController(userCourseUsecase)

	summaryUsecase := _summaryUsecase.NewSummaryUsecase(timeout, courseUsecase, userUsecase)
	summaryController := _summaryController.NewSummaryController(summaryUsecase)

	mysqlSpesializationRepository := _driversFectory.NewMysqlSpesializationRepository(conn)
	spesializationUsecae := _spesializationUsecase.NewSpesializationUsecase(mysqlSpesializationRepository, timeout)
	spesializationController := _spesializationsController.NewSpesializationController(spesializationUsecae)

	mysqlModulRepository := _driversFectory.NewMysqlModulRepository(conn)
	modulUsecase := _modulsUsecase.NewModulUsecase(mysqlModulRepository, courseUsecase, timeout)
	modulCtrl := _modulController.NewModulsController(modulUsecase)

	routesInit := routes.ControllerList{
		UserController:           *userCtrl,
		CourseController:         *couserCtrl,
		JWTMiddleware:            configJwt.Init(),
		UserCourseController:     *userCourseController,
		ModulController:          *modulCtrl,
		SummaryController:        *summaryController,
		SpesializationController: *spesializationController,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
