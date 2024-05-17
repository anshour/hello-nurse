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

func (dbase *V1User) NurseRegister(c echo.Context) (err error) {
	var req entities.NurseRegister

	if err := validator.BindValidate(c, &req); err != nil {
		println("error nih")
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	result := validator.ValidateNip(req.Nip, constants.NipNurse)

	if result != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: result.Error(),
		})
	}

	user := userUsecase.New(userRepository.New(dbase.DB))
	println("masuk sini ", req.Name)
	resp, err := user.CreateNurse(&entities.NurseRegisterParams{
		Nip:          req.Nip,
		Name:         req.Name,
		IdentityCard: req.IdentityImage,
		Role:         constants.ROLE_NURSE,
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
		Message: "Nurse registered successfully",
		Data:    resp,
	})

}
