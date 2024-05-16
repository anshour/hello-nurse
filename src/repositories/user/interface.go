package userRepository

import (
	"database/sql"
	entitiesLogin "hello-nurse/src/entities/user"
	entities "hello-nurse/src/entities/user/it"
)

type controllerUser struct {
	DB *sql.DB
}

type UserRepositories interface {
	Create(*entities.ITRegisterParams) (*entities.ITRegisterResponse, error)
	Login(*entitiesLogin.LoginParams) (*entitiesLogin.LoginResponse, error)
	List(*entities.UserListFilter) ([]*entities.UserListResponse, error)
}

func New(db *sql.DB) UserRepositories {
	return &controllerUser{DB: db}
}
