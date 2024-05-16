package medicalUsecase

import (
	entities "hello-nurse/src/entities/medical"
)

func (i *sMedicalUsecase) CreatePatient(p *entities.PatientRegisterParams) (*entities.PatientRegisterResult, error) {

	data, err := i.medicalRepository.CreatePatient(&entities.PatientRegisterParams{
		IdentityNumber: p.IdentityNumber,
		Name:           p.Name,
		PhoneNumber:    p.PhoneNumber,
		BirthDate:      p.BirthDate,
		Gender:         p.Gender,
		IdentityCard:   p.IdentityCard,
	})

	if err != nil {
		return nil, err
	}

	return &entities.PatientRegisterResult{
		UserId: data.UserId,
	}, nil
}
