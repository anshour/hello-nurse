package medical

type RecordRegister struct {
	IdentityNumber string `json:"identity_number" validate:"required,min=13,max=13"`
	Symptoms       string `json:"symptoms" validate:"required,min=5,max=255"`
	Medications    string `json:"medications" validate:"required,min=5,max=255"`
}
