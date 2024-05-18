package userRepository

import (
	entities "hello-nurse/src/entities/user"
	"log"
)

func (i *controllerUser) Create(params *entities.ITRegisterParams) (*entities.ITRegisterResponse, error) {
	var UserId string
	println("params.Role: ", params.Role)
	err := i.DB.QueryRow(`INSERT INTO users (nip, name, password, role, has_access) 
    VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		params.Nip, params.Name, params.Password, params.Role, true,
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
