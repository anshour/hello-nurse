package medicalController

import (
	entities "hello-nurse/src/entities/medical"
	record "hello-nurse/src/http/models/medical/record"
	medicalRepository "hello-nurse/src/repositories/medical"
	medicalUsecase "hello-nurse/src/usecase/medical"
	"hello-nurse/src/utils/validator"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (dbase *V1Medical) RecordRegister(c echo.Context) (err error) {

	var req record.RecordRegister

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}
	patient := medicalUsecase.New(medicalRepository.New(dbase.DB))

	resp, err := patient.CreateRecord(&entities.RecordRegisterParams{
		IdentityNumber: req.IdentityNumber,
		Symptoms:       req.Symptoms,
		Medications:    req.Medications,
	})

	if err != nil {

		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" {
				return c.JSON(http.StatusConflict, ErrorResponse{Message: "Indentity Number already exist"})
			}

			return c.JSON(http.StatusConflict, ErrorResponse{
				Message: err.Detail,
			})
		}

		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Success",
		Data:    resp,
	})

}
