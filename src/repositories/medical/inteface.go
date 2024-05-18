package medicalRepository

import (
	"database/sql"
	entities "hello-nurse/src/entities/medical"
)

type controllerMedical struct {
	DB *sql.DB
}

type MedicalRepositories interface {
	CreatePatient(*entities.PatientRegisterParams) (*entities.PatientRegisterResponse, error)
	ListPatient(*entities.PatientListFilter) ([]*entities.PatientListResult, error)
	CreateRecords(*entities.RecordRegisterParams) (*entities.RecordRegisterResponse, error)
}

func New(db *sql.DB) MedicalRepositories {
	return &controllerMedical{DB: db}
}
