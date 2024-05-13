package v1routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (i *V1Routes) MountNurse() {
	g := i.Echo.Group("/user")

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from it register")
	})
	g.POST("/nurse/register", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from it register")
	})
	g.PUT("/nurse/:userId", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from it register")
	})
	g.DELETE("/nurse/:userId", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from it register")
	})
	g.POST("/nurse/:userId/access", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from it register")
	})

}
