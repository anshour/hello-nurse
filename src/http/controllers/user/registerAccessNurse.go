package userController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (i *V1User) NurseAccess(c echo.Context) (err error) {

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User registered successfully",
		Data:    "",
	})

}
