package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"profcourse/app/routes"
	_userUsecase "profcourse/business/users"
	_userController "profcourse/controllers/users"
	_driversFectory "profcourse/drivers"
	_userMysqlRepo "profcourse/drivers/databases/users"
	_dbDriver "profcourse/drivers/mysql"
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

	conn := configDB.InitialDB()
	DbMigration(conn)

	e := echo.New()

	timeout := time.Duration(viper.GetInt("context.timeout")) * time.Second

	mysqlUserRepository := _driversFectory.NewMysqlUserRepository(conn)
	userUsecase := _userUsecase.NewUserUsecase(mysqlUserRepository, timeout)
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := routes.ControllerList{UserController: *userCtrl}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
