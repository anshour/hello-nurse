package medicalUsecase

import (
	entities "hello-nurse/src/entities/medical"
)

func (i *sMedicalUsecase) CreateRecord(p *entities.RecordRegisterParams) (*entities.RecordRegisterResponse, error) {

	data, err := i.medicalRepository.CreateRecords(p)

	if err != nil {
		return nil, err
	}

	return &entities.RecordRegisterResponse{
		Id: data.Id,
	}, nil
}
