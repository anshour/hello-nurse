package userUsecase

import (
	entities "hello-nurse/src/entities/user/it"
	"hello-nurse/src/utils/jwt"
	"hello-nurse/src/utils/password"
)

func (i *sUserUsecase) CreateUser(p *entities.ParamsITRegister) (*entities.ResultITRegister, error) {

	hashedPassword := password.Hash(p.Password)

	data, err := i.userRepository.Create(&entities.ParamsITRegister{
		Nip:      p.Nip,
		Name:     p.Name,
		Password: hashedPassword,
		Role:     "it",
	})

	if err != nil {
		return nil, err
	}

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: data.Id,
		Role:   "it",
	})

	return &entities.ResultITRegister{
		UserId:      data.Id,
		Nip:         p.Nip,
		Name:        p.Name,
		AccessToken: accessToken,
	}, nil
}
