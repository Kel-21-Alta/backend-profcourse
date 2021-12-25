package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"profcourse/app/middlewares"
	"profcourse/app/routes"
	_coursesUsecase "profcourse/business/courses"
	_userUsecase "profcourse/business/users"
	"profcourse/controllers/courses"
	_userController "profcourse/controllers/users"
	_driversFectory "profcourse/drivers"
	_userMysqlRepo "profcourse/drivers/databases/users"
	_dbDriver "profcourse/drivers/mysql"
	"profcourse/drivers/thirdparties/smtp"
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
	err = db.AutoMigrate(&_userMysqlRepo.User{})
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

	congfigSmtp := smtp.SmtpEmail{
		ConfigSmtpHost:     viper.GetString("smtp.host"),
		ConfigSmtpPort:     viper.GetInt("smtp.port"),
		ConfigSenderName:   viper.GetString("smtp.name"),
		ConfigAuthEmail:    viper.GetString("smtp.email"),
		ConfigAuthPassword: viper.GetString("smtp.password"),
	}

	configJwt := middlewares.ConfigJwt{
		SecretJwt:       viper.GetString("jwt.secret"),
		ExpiredDuration: viper.GetInt("jwt.expired"),
	}

	conn := configDB.InitialDB()
	DbMigration(conn)

	e := echo.New()

	timeout := time.Duration(viper.GetInt("context.timeout")) * time.Second

	smtpRepository := _driversFectory.NewSmtpRepository(congfigSmtp)
	localRepository := _driversFectory.NewLocalRepository()

	mysqlUserRepository := _driversFectory.NewMysqlUserRepository(conn)
	userUsecase := _userUsecase.NewUserUsecase(mysqlUserRepository, timeout, smtpRepository, configJwt)
	userCtrl := _userController.NewUserController(userUsecase)

	mysqlCourseRepository := _driversFectory.NewMysqlCourseRepository(conn)
	courseUsecase := _coursesUsecase.NewCourseUseCase(mysqlCourseRepository, timeout, localRepository)
	couserCtrl := courses.NewCourseController(courseUsecase)

	routesInit := routes.ControllerList{UserController: *userCtrl, CourseController: *couserCtrl, JWTMiddleware: configJwt.Init()}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
