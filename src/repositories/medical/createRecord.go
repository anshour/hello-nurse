package medicalRepository

import (
	"errors"
	entities "hello-nurse/src/entities/medical"
	"log"
)

func (dbase *controllerMedical) CreateRecords(params *entities.RecordRegisterParams) (*entities.RecordRegisterResponse, error) {
	var UserId string

	rows, err := dbase.DB.Exec(`INSERT INTO medical_records (identity_number, symptoms, medications) 
    SELECT $1, $2, $3
	WHERE EXISTS (SELECT 1 FROM patients WHERE identity_number = $4) RETURNING id`,
		params.IdentityNumber, params.Symptoms,
		params.Medications, params.IdentityNumber,
	)

	if err != nil {
		println(err)
		log.Printf("Error inserting records: %s", err)
		return nil, err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		log.Printf("Error checking row affected: %s", err)
		return nil, err
	}

	if rowsAffected == 0 {
		log.Printf("Error no row affected: %s", err)
		return nil, errors.New("identity number is not exist")

	}

	user := &entities.RecordRegisterResponse{
		UserId: UserId,
	}

	return user, nil
}
