package userUsecase

import (
	entities "hello-nurse/src/entities/user"
)

func (i *sUserUsecase) DeleteNurse(p *entities.NurseDeleteParams) error {

	err := i.userRepository.DeleteNurse(&entities.NurseDeleteParams{
		UserId: p.UserId,
	})

	if err != nil {
		return err
	}

	return nil
}
