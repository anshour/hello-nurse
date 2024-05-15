package userRepository

import (
	"database/sql"
	entities "hello-nurse/src/entities/user/it"
)

type controllerUser struct {
	DB *sql.DB
}

type UserRepositories interface {
	Create(*entities.ITRegisterParams) (*entities.ITRegisterResponse, error)
	Login(*entities.ITLoginParams) (*entities.ITLoginResponse, error)
}

func New(db *sql.DB) UserRepositories {
	return &controllerUser{DB: db}
}
