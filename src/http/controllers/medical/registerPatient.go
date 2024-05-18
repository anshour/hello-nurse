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

func (dbase *V1Medical) PatientRegister(c echo.Context) (err error) {
	var req entities.PatientRegister

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	if req.PhoneNumber != "" {
		err = validator.ValidatePhoneNumber(req.PhoneNumber)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Message: err.Error(),
			})
		}
	}

	if req.Gender != "" {
		err = validator.ValidateGender(req.Gender)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Message: err.Error(),
			})
		}
	}

	err = validator.ValidateUrl(req.IdentityCard)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}

	patient := medicalUsecase.New(medicalRepository.New(dbase.DB))

	resp, err := patient.CreatePatient(&entities.PatientRegisterParams{
		IdentityNumber: req.IdentityNumber,
		BirthDate:      req.BirthDate,
		Gender:         req.Gender,
		Name:           req.Name,
		PhoneNumber:    req.PhoneNumber,
		IdentityCard:   req.IdentityCard,
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
		Message: "Patient Registered Successfully",
		Data:    resp,
	})

}
