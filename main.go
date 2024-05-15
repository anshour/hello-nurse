package main

import (
	"fmt"
	"hello-nurse/src/db"
	v1routes "hello-nurse/src/http/routes"
	"hello-nurse/src/utils/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"net/http"
)

func main() {
	dbConnection, err := db.CreateConnection()
	if err != nil {
		fmt.Println("Error creating database connection:", err)
		return
	}

	defer func() {
		if err := dbConnection.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Validator = validator.CustomValidator

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1 := &v1routes.V1Routes{
		Echo: e.Group("/v1"),
		DB:   dbConnection,
	}

	v1.MountAll()

	e.Logger.Fatal(e.Start(":8080"))
}
