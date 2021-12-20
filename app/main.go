package main

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"profcourse/app/routes"
	_userUsecase "profcourse/business/users"
	_userController "profcourse/controllers/users"
	_dbDriver "profcourse/drivers/mysql"
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

func main() {
	e := echo.New()

	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}

	configDB.InitialDB()

	userUsecase := _userUsecase.NewUserUsecase()
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := routes.ControllerList{UserController: *userCtrl}

	routesInit.RouteRegister(e)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
