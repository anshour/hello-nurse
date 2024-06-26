package userUsecase

import (
	"hello-nurse/src/constants"
	entities "hello-nurse/src/entities/user"
	"hello-nurse/src/utils/jwt"
	"hello-nurse/src/utils/password"
	"log"
)

func (i *sUserUsecase) LoginUser(p *entities.LoginParams) (*entities.LoginResult, error) {

	hashedPassword := password.Hash(p.Password)

	data, err := i.userRepository.Login(&entities.LoginParams{
		Nip:      p.Nip,
		Password: hashedPassword,
	})

	if err != nil {
		return nil, err
	}

	err = password.Verify(data.Password, p.Password)
	if err != nil {
		log.Printf("Error Password verify: %s", err)
		return nil, constants.ErrNotFound

	}

	accessToken := jwt.Generate(&jwt.TokenPayload{
		UserId: data.Id,
		Role:   data.Role,
	})

	return &entities.LoginResult{
		UserId:      data.Id,
		Name:        data.Name,
		AccessToken: accessToken,
		Nip:         data.Nip,
	}, nil
}
