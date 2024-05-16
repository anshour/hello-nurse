package userController

import (
	"fmt"
	"hello-nurse/src/constants"
	entities "hello-nurse/src/entities/user"
	user "hello-nurse/src/http/models/user/it"
	userRepository "hello-nurse/src/repositories/user"
	userUsecase "hello-nurse/src/usecase/user"
	"hello-nurse/src/utils/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (dbase *V1User) ITRegister(c echo.Context) (err error) {
	var req user.ITRegister

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	nipStr := strconv.FormatInt(req.Nip, 10)
	if nipStr[0:3] != constants.NipIT {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintln("Nip must start with $1", constants.NipIT),
		})
	}

	userIT := userUsecase.New(userRepository.New(dbase.DB))

	resp, err := userIT.CreateUser(&entities.ITRegisterParams{
		Nip:      req.Nip,
		Name:     req.Name,
		Password: req.Password,
	})

	if err != nil {

		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" {
				return c.JSON(http.StatusConflict, ErrorResponse{Message: "Nip already exist"})
			}

			return c.JSON(http.StatusConflict, ErrorResponse{
				Message: err.Detail,
			})
		}

		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User registered successfully",
		Data:    resp,
	})

}
