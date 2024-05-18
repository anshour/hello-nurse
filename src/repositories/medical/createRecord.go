package medicalRepository

import (
	entities "hello-nurse/src/entities/medical"
	"log"
)

func (dbase *controllerMedical) CreateRecords(params *entities.RecordRegisterParams) (*entities.RecordRegisterResponse, error) {
	var RecordId string

	err := dbase.DB.QueryRow(`INSERT INTO medical_records (identity_number, user_id, symptoms, medications) 
		VALUES ($1, $2, $3, $4)  RETURNING id`,
		params.IdentityNumber,
		params.UserId,
		params.Symptoms,
		params.Medications,
	).Scan(&RecordId)

	if err != nil {
		println(err)
		log.Printf("Error inserting records: %s", err)
		return nil, err
	}

	user := &entities.RecordRegisterResponse{
		Id: RecordId,
	}

	return user, nil
}
