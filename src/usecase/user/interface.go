package userUsecase

import (
	entities "hello-nurse/src/entities/user/it"
	userRepository "hello-nurse/src/repositories"
)

type sUserUsecase struct {
	userRepository userRepository.UserRepositories
}

type UserUsecase interface {
	CreateUser(*entities.ParamsITRegister) (*entities.ResultITRegister, error)
	// Login(*ParamsLogin) (*ResultLogin, error)
}

func New(
	userRepository userRepository.UserRepositories,
) UserUsecase {
	return &sUserUsecase{
		userRepository: userRepository,
	}
}
