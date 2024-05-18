package userController

import (
	"hello-nurse/src/constants"
	entities "hello-nurse/src/entities/user"
	userRepository "hello-nurse/src/repositories/user"
	userUsecase "hello-nurse/src/usecase/user"
	"hello-nurse/src/utils/validator"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// TODO: MAKE IT LOGIN AND NURSE LOGIN AS ONE FUNCTION

func (dbase *V1User) ITLogin(c echo.Context) (err error) {
	var req entities.UserLogin

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	check := validator.ValidateNip(req.Nip, constants.NipIT)

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

	if result == nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: "Nip not exist"})

	}

	if err != nil {
		println("errrrr:", err.Error())

		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" {
				return c.JSON(http.StatusConflict, ErrorResponse{Message: "Nip already exist"})
			}

			return c.JSON(http.StatusConflict, ErrorResponse{
				Message: err.Detail,
			})
		}

		return c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Success",
		Data:    result,
	})

}
