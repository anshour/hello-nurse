package userRepository

import (
	"database/sql"
	entities "hello-nurse/src/entities/user"
)

type controllerUser struct {
	DB *sql.DB
}

type UserRepositories interface {
	Create(*entities.ITRegisterParams) (*entities.ITRegisterResponse, error)
	Login(*entities.LoginParams) (*entities.LoginResponse, error)
	CreateNurse(*entities.NurseRegisterParams) (*entities.NurseRegisterResponse, error)
	List(*entities.UserListFilter) ([]*entities.UserListResponse, error)
}

func New(db *sql.DB) UserRepositories {
	return &controllerUser{DB: db}
}
