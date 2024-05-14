package medicalController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (i *V1Medical) PatientRegister(c echo.Context) (err error) {

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Success",
		Data:    "",
	})

}
