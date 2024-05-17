package userUsecase

import (
	entities "hello-nurse/src/entities/user"
	"hello-nurse/src/utils/password"
)

func (i *sUserUsecase) CreateNurseAccess(p *entities.NurseRegisterAccess) error {

	hashedPassword := password.Hash(p.Password)

	err := i.userRepository.CreateNurseAccess(&entities.NurseRegisterAccess{
		Password: hashedPassword,
		UserId:   p.UserId,
	})

	if err != nil {
		return err
	}

	return nil
}
