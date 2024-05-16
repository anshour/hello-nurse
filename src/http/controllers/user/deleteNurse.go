package userController

import (
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

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Nurse Delete successfully",
		Data:    "",
	})

}
