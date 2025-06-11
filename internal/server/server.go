package server

import (
	"GinEchoCrud/internal/controller"
	"GinEchoCrud/internal/middleware"
	"GinEchoCrud/internal/repository"
	"GinEchoCrud/internal/service"
	"github.com/labstack/echo/v4"
)

func InitNewServer() {
	e := echo.New()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService)

	e.POST("/login", authController.Login)

	protected := e.Group("")
	protected.Use(middleware.JWTMiddleware)
	protected.GET("/users", userController.GetAllUsers)

	e.Logger.Fatal(e.Start(":1323"))
}
