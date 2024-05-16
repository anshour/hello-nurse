package userUsecase

import (
	entities "hello-nurse/src/entities/user"
	userRepository "hello-nurse/src/repositories/user"
)

type sUserUsecase struct {
	userRepository userRepository.UserRepositories
}

type UserUsecase interface {
	CreateUser(*entities.ITRegisterParams) (*entities.ITRegisterResult, error)
	LoginUser(*entities.LoginParams) (*entities.LoginResult, error)
	CreateNurse(*entities.NurseRegisterParams) (*entities.NurseRegisterResult, error)
	CreateNurseAccess(*entities.NurseRegisterAccess) error
	DeleteNurse(*entities.NurseDeleteParams) error
	EditNurse(*entities.NurseEditParams) error
	ListUser(*entities.UserListFilter) ([]*entities.UserListResponse, error)
}

func New(
	userRepository userRepository.UserRepositories,
) UserUsecase {
	return &sUserUsecase{
		userRepository: userRepository,
	}
}
