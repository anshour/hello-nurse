package medicalUsecase

import (
	entities "hello-nurse/src/entities/medical"
)

func (i *sMedicalUsecase) CreateRecord(p *entities.RecordRegisterParams) (*entities.RecordRegisterResult, error) {

	data, err := i.medicalRepository.CreateRecords(&entities.RecordRegisterParams{
		IdentityNumber: p.IdentityNumber,
		Symptoms:       p.Symptoms,
		Medications:    p.Medications,
	})

	if err != nil {
		return nil, err
	}

	return &entities.RecordRegisterResult{
		UserId: data.UserId,
	}, nil
}
