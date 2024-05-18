package medicalRepository

import (
	entities "hello-nurse/src/entities/medical"
	"log"
)

func (dbase *controllerMedical) CreatePatient(params *entities.PatientRegisterParams) (*entities.PatientRegisterResponse, error) {
	var UserId string
	err := dbase.DB.QueryRow(`INSERT INTO patients (identity_number, phone_number, name, birth_date, gender, identity_image) 
    VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		params.IdentityNumber, params.PhoneNumber,
		params.Name, params.BirthDate, params.Gender, params.IdentityCard,
	).Scan(&UserId)

	if err != nil {
		log.Printf("Error inserting patient: %s", err)
		return nil, err
	}

	user := &entities.PatientRegisterResponse{
		UserId: UserId,
	}

	return user, nil
}
