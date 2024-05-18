package userController

import (
	"hello-nurse/src/constants"
	entities "hello-nurse/src/entities/user"
	userRepository "hello-nurse/src/repositories/user"
	userUsecase "hello-nurse/src/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (dbase *V1User) NurseDelete(c echo.Context) (err error) {
	userIdParams := c.Param("userId")

	if userIdParams == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'user id'",
		})
	}
	user := userUsecase.New(userRepository.New(dbase.DB))

	err = user.DeleteNurse(&entities.NurseDeleteParams{
		UserId: userIdParams,
	})

	if err == constants.ErrNotFound {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}
	if err == constants.ErrInternalServer {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Nurse Delete successfully",
		Data:    "",
	})

}
