package medicalController

import (
	entities "hello-nurse/src/entities/medical"
	medicalRepository "hello-nurse/src/repositories/medical"
	medicalUsecase "hello-nurse/src/usecase/medical"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (dbase *V1Medical) PatientList(c echo.Context) (err error) {
	filters := &entities.PatientListFilter{}

	if limitStr := c.QueryParam("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'limit'",
			})
		}
		filters.Limit = limit
	}
	if offsetStr := c.QueryParam("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'offset'",
			})
		}
		filters.Offset = offset
	}

	if id := c.QueryParam("identityNumber"); id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'identity number'",
			})
		}
		filters.IdentityNumber = idInt
	}

	if name := c.QueryParam("name"); name != "" {
		filters.Name = name
	}

	if phone := c.QueryParam("phoneNumber"); phone != "" {
		filters.PhoneNumber = phone
	}

	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		filters.CreatedAt = createdAt
	}

	patient := medicalUsecase.New(medicalRepository.New(dbase.DB))
	patients, err := patient.ListPatient(filters)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Success",
		Data:    patients,
	})

}
