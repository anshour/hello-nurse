package userRepository

import (
	entities "hello-nurse/src/entities/user"
	"log"
)

func (i *controllerUser) CreateNurse(params *entities.NurseRegisterParams) (*entities.NurseRegisterResponse, error) {
	var UserId string

	println("masuk repo: ", params.IdentityCard)
	err := i.DB.QueryRow(`INSERT INTO users (nip, name, role, identity_image) 
    VALUES ($1, $2, $3, $4) RETURNING id`,
		params.Nip, params.Name,
		params.Role, params.IdentityCard,
	).Scan(&UserId)

	if err != nil {
		log.Printf("Error inserting user: %s", err)
		return nil, err
	}

	user := &entities.NurseRegisterResponse{
		UserId: UserId,
	}

	return user, nil
}
