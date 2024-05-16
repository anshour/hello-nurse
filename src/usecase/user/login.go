package userUsecase

import (
	entities "hello-nurse/src/entities/user/it"
	"hello-nurse/src/utils/jwt"
	"hello-nurse/src/utils/password"
	"log"
)

func (i *sUserUsecase) LoginUser(p *entitites.LoginParams) (*entities.LoginResult, error) {

	hashedPassword := password.Hash(p.Password)

	data, err := i.userRepository.Login(&entitites.LoginParams{
		Nip:      p.Nip,
		Password: hashedPassword,
	})

	if err != nil {
		return nil, err
	}

	errs := password.Verify(data.Password, p.Password)
	if errs != nil {
		log.Printf("Error Password verify: %s", err)
		return nil, err

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
