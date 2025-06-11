package controller

import (
	"GinEchoCrud/dto"
	"GinEchoCrud/internal/helpers"
	"GinEchoCrud/internal/interfaces"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	userService interfaces.UserServiceInterface
}

func NewAuthController(userService interfaces.UserServiceInterface) *AuthController {
	return &AuthController{userService: userService}
}

func (a *AuthController) Login(ctx echo.Context) error {
	var login dto.LoginRequest
	if err := ctx.Bind(&login); err != nil {
		return helpers.ResponseError(ctx, http.StatusBadRequest, "invalid data")
	}

	user, err := a.userService.Login(login.Email, login.Password)
	if err != nil {
		return helpers.ResponseError(ctx, http.StatusUnauthorized, err.Error())
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return helpers.ResponseError(ctx, http.StatusInternalServerError, "error while creating token")
	}

	return helpers.ResponseSuccess(ctx, map[string]string{"token": token})
}
