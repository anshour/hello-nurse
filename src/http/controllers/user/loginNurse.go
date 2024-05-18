package userController

import (
	"hello-nurse/src/constants"
	entities "hello-nurse/src/entities/user"
	userRepository "hello-nurse/src/repositories/user"
	userUsecase "hello-nurse/src/usecase/user"
	"hello-nurse/src/utils/validator"

	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: MAKE IT LOGIN AND NURSE LOGIN AS ONE FUNCTION

func (dbase *V1User) NurseLogin(c echo.Context) (err error) {
	var req entities.UserLogin

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	check := validator.ValidateNip(req.Nip, constants.NipNurse)
	if check != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: check.Error(),
		})
	}

	data := userUsecase.New(userRepository.New(dbase.DB))
	result, err := data.LoginUser(&entities.LoginParams{
		Nip:      req.Nip,
		Password: req.Password,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Success",
		Data:    result,
	})

}
