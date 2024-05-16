package medicalUsecase

import (
	entities "hello-nurse/src/entities/medical"
	medicalRepository "hello-nurse/src/repositories/medical"
)

type sMedicalUsecase struct {
	medicalRepository medicalRepository.MedicalRepositories
}

type MedicalUsecase interface {
	CreatePatient(*entities.PatientRegisterParams) (*entities.PatientRegisterResult, error)
	CreateRecord(*entities.RecordRegisterParams) (*entities.RecordRegisterResult, error)
}

func New(
	medicalRepository medicalRepository.MedicalRepositories,
) MedicalUsecase {
	return &sMedicalUsecase{
		medicalRepository: medicalRepository,
	}
}
