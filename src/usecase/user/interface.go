package userUsecase

import (
	entities "hello-nurse/src/entities/user/it"
	userRepository "hello-nurse/src/repositories/user"
)

type sUserUsecase struct {
	userRepository userRepository.UserRepositories
}

type UserUsecase interface {
	CreateUser(*entities.ITRegisterParams) (*entities.ITRegisterResult, error)
	LoginUser(*entities.ITLoginParams) (*entities.ITLoginResult, error)
}

func New(
	userRepository userRepository.UserRepositories,
) UserUsecase {
	return &sUserUsecase{
		userRepository: userRepository,
	}
}
