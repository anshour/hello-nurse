package medicalRepository

import (
	entities "hello-nurse/src/entities/medical"
	"log"
)

func (dbase *controllerMedical) CreateRecords(params *entities.RecordRegisterParams) (*entities.RecordRegisterResponse, error) {
	var UserId string
	err := dbase.DB.QueryRow(`INSERT INTO medical_records (patient_id, symptoms, medications) 
    VALUES ($1, $2, $3) RETURNING id`,
		params.IdentityNumber, params.Symptoms,
		params.Medications,
	).Scan(&UserId)

	if err != nil {
		log.Printf("Error inserting records: %s", err)
		return nil, err
	}

	user := &entities.RecordRegisterResponse{
		UserId: UserId,
	}

	return user, nil
}
