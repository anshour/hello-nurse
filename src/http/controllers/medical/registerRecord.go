package medicalController

import (
	"hello-nurse/src/constants"
	entities "hello-nurse/src/entities/medical"
	medicalRepository "hello-nurse/src/repositories/medical"
	medicalUsecase "hello-nurse/src/usecase/medical"
	"hello-nurse/src/utils/validator"
	"net/http"
	"strconv"

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

	identityNumberStr := strconv.Itoa(req.IdentityNumber)
	if len(identityNumberStr) != 16 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Identity number must 16 digit",
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
			if err.Code == constants.UniqueViolationExistData {
				return c.JSON(http.StatusConflict, ErrorResponse{Message: "Indentity Number already exist"})
			}

			if err.Code == constants.UniqueViolationNotExistData {
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
