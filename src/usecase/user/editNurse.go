package userUsecase

import (
	entities "hello-nurse/src/entities/user"
)

func (i *sUserUsecase) EditNurse(p *entities.NurseEditParams) error {

	err := i.userRepository.EditNurse(&entities.NurseEditParams{
		UserId: p.UserId,
		Nip:    p.Nip,
		Name:   p.Name,
	})

	if err != nil {
		return err
	}

	return nil
}
