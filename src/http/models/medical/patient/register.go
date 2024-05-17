package medical

type PatientRegister struct {
	IdentityNumber string `json:"identity_number" validate:"required,min=13,max=13"`
	PhoneNumber    string `json:"phone_number" validate:"required,min=5,max=16"`
	Name           string `json:"name" validate:"required,min=5,max=50"`
	BirtDate       string `json:"birth_date" validate:"required,min=2,max=70"`
	Gender         string `json:"gender" validate:"required,min=4,max=6"`
	IdentityCard   string `json:"identity_card" validate:"required,min=4,max=255"`
}
