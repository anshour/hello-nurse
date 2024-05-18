package medicalController

import (
	entities "hello-nurse/src/entities/medical"
	medicalRepository "hello-nurse/src/repositories/medical"
	medicalUsecase "hello-nurse/src/usecase/medical"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (dbase *V1Medical) RecordList(c echo.Context) (err error) {
	filters := &entities.RecordListFilter{}

	if id := c.QueryParam("identityDetail.identityNumber"); id != "" {
		identityNumber, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'offset'",
			})
		}
		filters.IdentityNumber = identityNumber
	}

	if nip := c.QueryParam("createdBy.nip"); nip != "" {
		filters.Nip = nip
	}

	if userId := c.QueryParam("createdBy.userId"); userId != "" {
		filters.UserId = userId
	}

	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		filters.CreatedAt = createdAt
	}

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

	patient := medicalUsecase.New(medicalRepository.New(dbase.DB))
	records, err := patient.ListRecord(filters)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Success",
		Data:    records,
	})

}
