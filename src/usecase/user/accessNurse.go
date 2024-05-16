package userUsecase

import (
	entities "hello-nurse/src/entities/user"
)

func (i *sUserUsecase) CreateNurseAccess(p *entities.NurseRegisterAccess) error {

	err := i.userRepository.CreateNurseAccess(&entities.NurseRegisterAccess{
		Password: p.Password,
	})

	if err != nil {
		return err
	}

	return nil
}
