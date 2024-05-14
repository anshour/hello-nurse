package userController

import (
	user "hello-nurse/src/http/models/user/it"
	"hello-nurse/src/utils/jwt"
	"hello-nurse/src/utils/password"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type ReturnData struct {
	accessToken string
	userId      string
	nip         string
	name        string
}

func (dbase *V1User) ITRegister(c echo.Context) (err error) {
	var req user.ITRegister

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	hashedPassword := password.Hash(req.Password)
	var nameRole string
	if role := req.Nip[0:2]; role == "615" {
		nameRole = "IT"
	} else {
		nameRole = "Nurse"
	}
	var UserId string
	err = dbase.DB.QueryRow("INSERT INTO users (nip, name, password, role) VALUES ($1, $2, $3, $4) RETURNING id", req.Nip, req.Name, hashedPassword, nameRole).Scan(&UserId)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return c.JSON(http.StatusConflict, ErrorResponse{
				Message: err.Detail,
			})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: UserId,
	})

	data := ReturnData{accessToken: accessToken, userId: UserId, nip: req.Nip, name: req.Name}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User registered successfully",
		Data:    data,
	})

}
