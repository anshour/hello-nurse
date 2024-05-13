package v1routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (i *V1Routes) MountUser() {
	g := i.Echo.Group("/user")

	g.POST("/it/register", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from it register")
	})

}
