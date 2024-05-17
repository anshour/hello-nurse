package userController

import (
	entities "hello-nurse/src/entities/user"
	userRepository "hello-nurse/src/repositories/user"
	userUsecase "hello-nurse/src/usecase/user"
	"hello-nurse/src/utils/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (dbase *V1User) NurseAccess(c echo.Context) (err error) {
	userIdParams := c.Param("userId")

	if userIdParams == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'user id'",
		})
	}

	var req entities.NurseRegisterAccess

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	data := userUsecase.New(userRepository.New(dbase.DB))
	err = data.CreateNurseAccess(&entities.NurseRegisterAccess{
		Password: req.Password,
		UserId:   userIdParams,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Nurse Access Granted",
		Data:    "",
	})

}
