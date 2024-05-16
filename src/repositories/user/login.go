package userRepository

import (
	entities "hello-nurse/src/entities/user/it"
	"log"
)

func (i *controllerUser) Login(params *entitites.LoginParams) (*entities.LoginResponse, error) {

	var userId string
	var nip int64
	var name string
	var role string
	var password string
	err := i.DB.QueryRow("SELECT id, nip, name, role, password FROM users WHERE nip = $1", params.Nip).Scan(&userId, &nip, &name, &role, &password)
	if err != nil {
		log.Printf("Error login user: %s", err)
		return nil, err
	}

	user := &entities.LoginResponse{
		Id:       userId,
		Nip:      nip,
		Name:     name,
		Password: password,
		Role:     role,
	}

	return user, nil
}
