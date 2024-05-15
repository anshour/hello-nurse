package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func OnlyRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole, _ := c.Get("role").(string)
			println("masuk role")
			if userRole != role {
				return c.JSON(http.StatusUnauthorized, ErrorResponse{
					Status:  false,
					Message: "Unauthorized",
				})
			}

			return next(c)
		}
	}
}
