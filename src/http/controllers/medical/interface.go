package medicalController

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1Medical struct {
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

type iV1Medical interface {
	PatientList(c echo.Context) error
	PatientRegister(c echo.Context) error
	RecordList(c echo.Context) error
	RecordRegister(c echo.Context) error
}

func New(v1Medical *V1Medical) iV1Medical {
	return v1Medical
}
