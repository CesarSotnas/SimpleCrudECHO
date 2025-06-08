package server

import (
	"GinEchoCrud/internal/controller"
	"GinEchoCrud/internal/repository"
	"GinEchoCrud/internal/service"
	"github.com/labstack/echo/v4"
)

func InitNewServer() {
	e := echo.New()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	e.GET("/users", userController.GetAllUsers)

	e.Logger.Fatal(e.Start(":1323"))
}
