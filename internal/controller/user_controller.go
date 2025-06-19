package controller

import (
	"GinEchoCrud/internal/helpers"
	"GinEchoCrud/internal/interfaces"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserController struct {
	userService interfaces.UserServiceInterface
}

func NewUserController(userService interfaces.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetAllUsers(ctx echo.Context) error {
	users, statusCode, err := c.userService.GetAllUsers()
	if err != nil {
		return helpers.ResponseError(ctx, statusCode, err.Error())

	}

	return helpers.ResponseSuccess(ctx, statusCode, users)
}

func (c *UserController) GetUsersByID(ctx echo.Context) error {
	idStr := ctx.Param("user_id")
	if idStr == "" {
		return helpers.ResponseError(ctx, http.StatusBadRequest, helpers.ErrMsgEmptyValue.Error())
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return helpers.ResponseError(ctx, http.StatusBadRequest, helpers.ErrMsgIdMustBeANumber.Error())
	}

	users, statusCode, errResponse := c.userService.GetUsersByID(id)
	if errResponse != nil {
		return helpers.ResponseError(ctx, statusCode, errResponse.Error())
	}

	return helpers.ResponseSuccess(ctx, statusCode, users)
}
