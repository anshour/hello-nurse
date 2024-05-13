package userController

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1User struct {
	DB *sql.DB
}

type iV1User interface {
	NurseRegister(c *echo.Context)
}

func New(v1User *V1User) iV1User {
	return v1User
}
