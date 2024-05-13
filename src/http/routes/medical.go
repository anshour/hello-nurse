package v1routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (i *V1Routes) MountMedical() {
	g := i.Echo.Group("/medical")

	g.GET("/patient", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from it register")
	})
	g.POST("/patient", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from it register")
	})
	g.GET("/record", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from it register")
	})
	g.POST("/record", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from it register")
	})

}
