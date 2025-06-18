package server

import (
	"GinEchoCrud/internal/controller"
	"GinEchoCrud/internal/middleware"
	"GinEchoCrud/internal/repository"
	"GinEchoCrud/internal/service"
	"github.com/labstack/echo/v4"
)

func InitNewServer(port string) {
	e := echo.New()

	// auth dependency injection
	authRepo := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	// user dependency injection
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	e.POST("/login", authController.Login)

	protected := e.Group("/api")
	protected.Use(middleware.JWTMiddleware)
	protected.GET("/users", userController.GetAllUsers)

	e.Logger.Fatal(e.Start(port))
}
