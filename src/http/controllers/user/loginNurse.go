package userController

import (
	"hello-nurse/src/http/models/user"
	"hello-nurse/src/utils/jwt"
	utilsPassword "hello-nurse/src/utils/password"
	"hello-nurse/src/utils/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (i *V1User) NurseLogin(c echo.Context) (err error) {
	var req user.UserLogin

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	nipStr := strconv.FormatInt(req.Nip, 10)
	if nipStr[0:3] != "303" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Nip must start with 303",
		})
	}

	var userId string
	var nip int64
	var name string
	var password string
	err = i.DB.QueryRow("SELECT id, nip, name, password FROM users WHERE nip = $1", req.Nip).Scan(&userId, &nip, &name, &password)

	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "user not found"})
		return
	}

	errs := utilsPassword.Verify(password, req.Password)
	if errs != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid password"})
		return

	}

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: userId,
	})

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Success",
		Data: LoginItResponse{
			AccessToken: accessToken,
			UserId:      userId,
			Nip:         nip,
			Name:        name,
		},
	})

}
