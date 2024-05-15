package userController

import (
	user "hello-nurse/src/http/models/user/it"
	"hello-nurse/src/utils/jwt"
	"hello-nurse/src/utils/password"
	"hello-nurse/src/utils/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type RegisterITResponse struct {
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
	Nip         int64  `json:"nip"`
	Name        string `json:"name"`
}

func (i *V1User) ITRegister(c echo.Context) (err error) {
	var req user.ITRegister

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

	role := "it"
	hashedPassword := password.Hash(req.Password)
	var UserId string
	err = i.DB.QueryRow(`INSERT INTO users (nip, name, password, role) 
    VALUES ($1, $2, $3, $4) RETURNING id`, req.Nip, req.Name, hashedPassword, role).Scan(&UserId)
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

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: UserId,
		Role:   "it",
	})

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "User registered successfully",
		Data: RegisterITResponse{
			AccessToken: accessToken,
			UserId:      UserId,
			Nip:         req.Nip,
			Name:        req.Name,
		},
	})

}
