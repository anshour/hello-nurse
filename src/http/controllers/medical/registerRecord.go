package medicalController

import (
	entities "hello-nurse/src/entities/medical"
	medicalRepository "hello-nurse/src/repositories/medical"
	medicalUsecase "hello-nurse/src/usecase/medical"
	"hello-nurse/src/utils/validator"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (dbase *V1Medical) RecordRegister(c echo.Context) (err error) {

	var req entities.RecordRegister
	UserId, _ := c.Get("userId").(string)

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}
	patient := medicalUsecase.New(medicalRepository.New(dbase.DB))

	resp, err := patient.CreateRecord(&entities.RecordRegisterParams{
		UserId:         UserId,
		IdentityNumber: req.IdentityNumber,
		Symptoms:       req.Symptoms,
		Medications:    req.Medications,
	})

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" {
				return c.JSON(http.StatusConflict, ErrorResponse{Message: "Indentity Number already exist"})
			}

			if err.Code == "23503" {
				return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Patient with identity number doesnt exist"})
			}

			return c.JSON(http.StatusConflict, ErrorResponse{
				Message: err.Detail,
			})
		}

		return c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Success",
		Data:    resp,
	})

}
