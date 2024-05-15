package userController

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1User struct {
	DB *sql.DB
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1User interface {
	ITRegister(c echo.Context) error
	ITLogin(c echo.Context) error
	NurseLogin(c echo.Context) error
	NurseRegister(c echo.Context) error
	NurseEdit(c echo.Context) error
	NurseDelete(c echo.Context) error
	NurseAccess(c echo.Context) error
	UserList(c echo.Context) error
}

func New(v1User *V1User) iV1User {
	return v1User
}
