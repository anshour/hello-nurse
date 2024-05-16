package medicalController

import (
	medical "hello-nurse/src/http/models/medical/patient"
	"hello-nurse/src/utils/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (dbase *V1Medical) PatientRegister(c echo.Context) (err error) {
	var req medical.PatientRegister

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}
	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Success",
		Data:    "",
	})

}
