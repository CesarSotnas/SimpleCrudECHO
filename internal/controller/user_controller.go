package controller

import (
	"GinEchoCrud/internal/helpers"
	"GinEchoCrud/internal/interfaces"
	"GinEchoCrud/internal/models"
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

func (c *UserController) CreateUser(ctx echo.Context) error {
	var requestUser models.User

	err := ctx.Bind(&requestUser)
	if err != nil {
		return helpers.ResponseError(ctx, helpers.StatusBadRequest, helpers.ErrMsgInvalidData.Error())
	}

	users, statusCode, errResponse := c.userService.CreateUser(requestUser)
	if statusCode >= 400 || errResponse != nil {
		return helpers.ResponseError(ctx, statusCode, errResponse.Error())
	}

	return helpers.ResponseSuccess(ctx, helpers.StatusCreated, users)
}
