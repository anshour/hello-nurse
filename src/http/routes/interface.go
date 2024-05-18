package v1routes

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1Routes struct {
	Echo *echo.Group
	DB   *sql.DB
}

func (i *V1Routes) MountAll() {
	i.MountUser()
	i.MountMedical()
	// i.MountUpload()
}
