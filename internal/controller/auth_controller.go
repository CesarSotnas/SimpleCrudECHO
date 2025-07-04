package controller

import (
	"GinEchoCrud/dto"
	"GinEchoCrud/internal/helpers"
	"GinEchoCrud/internal/interfaces"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService interfaces.AuthServiceInterface
}

func NewAuthController(authService interfaces.AuthServiceInterface) *AuthController {
	return &AuthController{authService: authService}
}

func (a *AuthController) Login(ctx echo.Context) error {
	var login dto.LoginRequest
	if err := ctx.Bind(&login); err != nil {
		return helpers.ResponseError(ctx, helpers.StatusBadRequest, helpers.ErrMsgInvalidData.Error())
	}

	user, err := a.authService.Login(login.Email, login.Password)
	if err != nil {
		return helpers.ResponseError(ctx, helpers.StatusUnauthorized, helpers.ErrMsgUnauthorized.Error())
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return helpers.ResponseError(ctx, helpers.StatusBadRequest, helpers.ErrMsgCreateTokenError.Error())
	}

	return helpers.ResponseSuccess(ctx, helpers.StatusOk, map[string]string{"token": token})
}
