package server

import (
	"GinEchoCrud/internal/controller"
	"GinEchoCrud/internal/middleware"
	"GinEchoCrud/internal/repository"
	"GinEchoCrud/internal/routes"
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

	// auth middleware
	e.POST(routes.Login, authController.Login)

	// routes
	protected := e.Group(routes.Prefix)
	protected.Use(middleware.JWTMiddleware)
	protected.GET(routes.GetUsers, userController.GetAllUsers)
	protected.GET(routes.GetUsersById, userController.GetUsersByID)
	protected.POST(routes.CreateUsers, userController.CreateUser)

	// server
	e.Logger.Fatal(e.Start(port))
}
