package controller

import (
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/utils"
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
		return utils.ResponseError(ctx, http.StatusBadRequest, err.Error())

	}

	return utils.ResponseSuccess(ctx, users)
}
