package userUsecase

import (
	entities "hello-nurse/src/entities/user"
)

func (i *sUserUsecase) CreateNurse(p *entities.NurseRegisterParams) (*entities.NurseRegisterResult, error) {

	data, err := i.userRepository.CreateNurse(&entities.NurseRegisterParams{
		Nip:  p.Nip,
		Name: p.Name,
		Role: "nurse",
	})

	if err != nil {
		return nil, err
	}

	return &entities.NurseRegisterResult{
		UserId: data.Id,
		Nip:    p.Nip,
		Name:   p.Name,
	}, nil
}
