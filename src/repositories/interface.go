package userRepository

import (
	"database/sql"
	entities "hello-nurse/src/entities/user/it"
)

type controllerRegisterIT struct {
	DB *sql.DB
}

type UserRepositories interface {
	Create(*entities.ParamsITRegister) (*entities.ITRegisterResponse, error)
}

func New(db *sql.DB) UserRepositories {
	return &controllerRegisterIT{DB: db}
}
