package middleware

import (
	"GinEchoCrud/internal/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return helpers.ResponseError(ctx, http.StatusUnauthorized, "Missing or malformed token")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := helpers.ParseToken(tokenStr)
		if err != nil {
			return helpers.ResponseError(ctx, http.StatusUnauthorized, "Invalid Token")
		}

		ctx.Set("user_id", claims["user_id"])
		return next(ctx)
	}
}
