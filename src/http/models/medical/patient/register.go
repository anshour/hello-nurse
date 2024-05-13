package medical

type PatientRegister struct {
	IdentityNumber string `json:"identity_number" binding:"required,min=13,max=13"`
	PhoneNumber    string `json:"phone_number" binding:"required,min=5,max=16"`
	Name           string `json:"name" binding:"required,min=5,max=50"`
	BirtDate       string `json:"birth_date" binding:"required,min=2,max=70"`
	Gender         string `json:"gender" binding:"required,min=4,max=6"`
}
