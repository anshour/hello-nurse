package userUsecase

import (
	entities "hello-nurse/src/entities/user"
	"hello-nurse/src/utils/jwt"
	"hello-nurse/src/utils/password"
)

func (i *sUserUsecase) CreateUser(p *entities.ITRegisterParams) (*entities.ITRegisterResult, error) {

	hashedPassword := password.Hash(p.Password)

	data, err := i.userRepository.Create(&entities.ITRegisterParams{
		Nip:      p.Nip,
		Name:     p.Name,
		Password: hashedPassword,
		Role:     p.Role,
	})

	if err != nil {
		return nil, err
	}

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: data.Id,
		Role:   p.Role,
	})

	return &entities.ITRegisterResult{
		UserId:      data.Id,
		Nip:         p.Nip,
		Name:        p.Name,
		AccessToken: accessToken,
	}, nil
}
