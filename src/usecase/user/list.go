package userUsecase

import (
	entities "hello-nurse/src/entities/user"
)

func (i *sUserUsecase) ListUser(filters *entities.UserListFilter) ([]*entities.UserListResponse, error) {
	allUsers, err := i.userRepository.List(filters)

	if err != nil {
		return nil, err
	}

	return allUsers, nil
}
