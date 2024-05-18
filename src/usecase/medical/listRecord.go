package medicalUsecase

import (
	entities "hello-nurse/src/entities/medical"
)

func (i *sMedicalUsecase) ListRecord(filters *entities.RecordListFilter) ([]*entities.RecordListResult, error) {

	records, err := i.medicalRepository.ListRecord(filters)

	if err != nil {
		return nil, err
	}

	return records, nil
}
