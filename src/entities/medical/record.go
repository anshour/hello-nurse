package entities

type RecordRegister struct {
	IdentityNumber string `json:"identityNumber" validate:"required,min=16,max=16"`
	Symptoms       string `json:"symptoms" validate:"required,min=5,max=255"`
	Medications    string `json:"medications" validate:"required,min=5,max=255"`
}

type RecordRegisterResponse struct {
	UserId string
}

type RecordRegisterResult struct {
	UserId string
}

type RecordRegisterParams struct {
	IdentityNumber string
	Symptoms       string
	Medications    string
}
