package medicalController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (dbase *V1Medical) RecordRegister(c echo.Context) (err error) {

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Success",
		Data:    "",
	})

}
