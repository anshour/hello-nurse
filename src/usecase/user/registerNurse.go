package userUsecase

import (
	entities "hello-nurse/src/entities/user"
)

func (i *sUserUsecase) CreateNurse(p *entities.NurseRegisterParams) (*entities.NurseRegisterResult, error) {

	data, err := i.userRepository.CreateNurse(&entities.NurseRegisterParams{
		Nip:          p.Nip,
		Name:         p.Name,
		Role:         p.Role,
		IdentityCard: p.IdentityCard,
	})

	if err != nil {
		println("error in usecase: ", err)
		return nil, err
	}

	return &entities.NurseRegisterResult{
		UserId: data.UserId,
		Nip:    p.Nip,
		Name:   p.Name,
	}, nil
}
