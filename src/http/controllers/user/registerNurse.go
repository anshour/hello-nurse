package userController

import (
	user "hello-nurse/src/http/models/user/nurse"
	"hello-nurse/src/utils/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type NurseRegisterResponse struct {
	UserId string `json:"userId"`
	Nip    int64  `json:"nip"`
	Name   string `json:"name"`
}

func (i *V1User) NurseRegister(c echo.Context) (err error) {
	var req user.NurseRegister

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

	role := "nurse"
	rawPassword := "xxx"
	var UserId string
	err = i.DB.QueryRow(`INSERT INTO users (nip, name, password, role, identity_image) 
    VALUES ($1, $2, $3, $4, $5) RETURNING id`, req.Nip, req.Name, rawPassword, role, req.IdentityImage).Scan(&UserId)
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
		Data: NurseRegisterResponse{
			UserId: UserId,
			Nip:    req.Nip,
			Name:   req.Name,
		},
	})

}
