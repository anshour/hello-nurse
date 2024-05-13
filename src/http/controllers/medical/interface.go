package medicalController

import (
	"database/sql"
)

type V1Medical struct {
	DB *sql.DB
}

type iV1Medical interface {
}

func New(v1Medical *V1Medical) iV1Medical {
	return v1Medical
}
