package controller

import (
	"GinEchoCrud/internal/interfaces"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userService interfaces.UserServiceInterface
}

func NewUserController(userService interfaces.UserRepositoryInterface) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetAllUsers(ctx echo.Context) error {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

	}

	return ctx.JSON(http.StatusOK, users)
}
