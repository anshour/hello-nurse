package userController

import (
	entities "hello-nurse/src/entities/user/it"
	userRepository "hello-nurse/src/repositories/user"
	userUsecase "hello-nurse/src/usecase/user"
	"hello-nurse/src/utils/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (dbase *V1User) UserList(c echo.Context) (err error) {

	filters := &entities.UserListFilter{}

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

	if id := c.QueryParam("userId"); id != "" {
		filters.Id = id
	}

	if name := c.QueryParam("name"); name != "" {
		filters.Name = name
	}
	if nip := c.QueryParam("nip"); nip != "" {
		nipInt, err := strconv.Atoi(nip)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'Nip'",
			})
		}
		filters.Nip = nipInt
	}

	if role := c.QueryParam("role"); role != "" {
		err = validator.ValidateCategory(role)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
		filters.Role = role
	}

	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		filters.CreatedAt = createdAt
	}

	userList := userUsecase.New(userRepository.New(dbase.DB))
	users, err := userList.ListUser(filters)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User list successfully",
		Data:    users,
	})

}
