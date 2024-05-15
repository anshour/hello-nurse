package userController

import (
	entities "hello-nurse/src/entities/user/it"
	"hello-nurse/src/http/models/user"
	userRepository "hello-nurse/src/repositories/user"
	userUsecase "hello-nurse/src/usecase/user"
	"hello-nurse/src/utils/validator"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginItResponse struct {
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
	Nip         int64  `json:"nip"`
	Name        string `json:"name"`
}

// TODO: MAKE IT LOGIN AND NURSE LOGIN AS ONE FUNCTION

func (dbase *V1User) ITLogin(c echo.Context) (err error) {
	var req user.UserLogin

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	nipStr := strconv.FormatInt(req.Nip, 10)
	if nipStr[0:3] != "615" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Nip must start with 615",
		})
	}

	data := userUsecase.New(userRepository.New(dbase.DB))
	result, err := data.LoginUser(&entities.ITLoginParams{
		Nip:      req.Nip,
		Password: req.Password,
	})

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Success",
		Data: LoginItResponse{
			AccessToken: result.AccessToken,
			UserId:      result.UserId,
			Name:        result.Name,
			Nip:         req.Nip,
		},
	})

}
