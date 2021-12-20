package main

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"profcourse/app/routes"
	_userUsecase "profcourse/business/users"
	_userController "profcourse/controllers/users"
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
	userUsecase := _userUsecase.NewUserUsecase()
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := routes.ControllerList{UserController: *userCtrl}

	routesInit.RouteRegister(e)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
