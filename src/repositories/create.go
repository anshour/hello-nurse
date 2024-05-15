package userRepository

import (
	entities "hello-nurse/src/entities/user/it"
	"log"
)

func (i *controllerRegisterIT) Create(params *entities.ParamsITRegister) (*entities.ITRegisterResponse, error) {
	var UserId string
	err := i.DB.QueryRow(`INSERT INTO users (nip, name, password, role) 
    VALUES ($1, $2, $3, $4) RETURNING id`,
		params.Nip, params.Name, params.Password,
		params.Role,
	).Scan(&UserId)

	if err != nil {
		log.Printf("Error inserting user: %s", err)
		return nil, err
	}

	user := &entities.ITRegisterResponse{
		Id: UserId,
	}

	return user, nil
}
