package medicalUsecase

import (
	entities "hello-nurse/src/entities/medical"
)

func (i *sMedicalUsecase) ListPatient(filters *entities.PatientListFilter) ([]*entities.PatientListResult, error) {

	patients, err := i.medicalRepository.ListPatient(filters)

	if err != nil {
		return nil, err
	}

	return patients, nil
}
