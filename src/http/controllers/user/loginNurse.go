package userController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (i *V1User) NurseLogin(c echo.Context) (err error) {

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Success",
		Data:    "",
	})

}
