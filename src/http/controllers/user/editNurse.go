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

func (dbase *V1User) NurseEdit(c echo.Context) (err error) {
	userIdParams := c.Param("userId")

	if userIdParams == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'user id'",
		})
	}

	var req entities.NurseEdit

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

	user := userUsecase.New(userRepository.New(dbase.DB))

	err = user.EditNurse(&entities.NurseEditParams{
		UserId: userIdParams,
		Nip:    req.Nip,
		Name:   req.Name,
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
		Message: "User Edit successfully",
	})

}
