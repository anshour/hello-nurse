package userRepository

import (
	entities "hello-nurse/src/entities/user"
	"log"
)

func (i *controllerUser) CreateNurse(params *entities.NurseRegisterParams) (*entities.NurseRegisterResponse, error) {
	var UserId string
	err := i.DB.QueryRow(`INSERT INTO users (nip, name, role) 
    VALUES ($1, $2, $3, $4) RETURNING id`,
		params.Nip, params.Name,
		params.Role,
	).Scan(&UserId)

	if err != nil {
		log.Printf("Error inserting user: %s", err)
		return nil, err
	}

	user := &entities.NurseRegisterResponse{
		Id: UserId,
	}

	return user, nil
}
