package userUsecase

import (
	entitiesLogin "hello-nurse/src/entities/user"
	entities "hello-nurse/src/entities/user/it"
	userRepository "hello-nurse/src/repositories/user"
)

type sUserUsecase struct {
	userRepository userRepository.UserRepositories
}

type UserUsecase interface {
	CreateUser(*entities.ITRegisterParams) (*entities.ITRegisterResult, error)
	LoginUser(*entitiesLogin.LoginParams) (*entitiesLogin.LoginResult, error)
	ListUser(*entities.UserListFilter) ([]*entities.UserListResponse, error)
}

func New(
	userRepository userRepository.UserRepositories,
) UserUsecase {
	return &sUserUsecase{
		userRepository: userRepository,
	}
}
