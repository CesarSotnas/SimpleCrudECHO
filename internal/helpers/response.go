package helpers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ResponseSuccess(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func ResponseError(ctx echo.Context, status int, message string) error {
	return ctx.JSON(status, map[string]interface{}{
		"success": false,
		"data":    message,
	})
}
