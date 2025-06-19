package helpers

import (
	"github.com/labstack/echo/v4"
)

func ResponseSuccess(ctx echo.Context, statusCode int, data interface{}) error {
	return ctx.JSON(statusCode, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func ResponseError(ctx echo.Context, statusCode int, message string) error {
	return ctx.JSON(statusCode, map[string]interface{}{
		"data":  nil,
		"error": message,
	})
}
