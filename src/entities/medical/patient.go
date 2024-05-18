package entities

type PatientRegister struct {
	IdentityNumber int    `json:"identityNumber" validate:"required"`
	PhoneNumber    string `json:"phoneNumber" validate:"required,min=10,max=15"`
	Name           string `json:"name" validate:"required,min=5,max=50"`
	BirthDate      string `json:"birthDate" validate:"required,min=2,max=70"`
	Gender         string `json:"gender" validate:"required,min=4,max=6"`
	IdentityCard   string `json:"identityCardScanImg" validate:"required,min=4,max=255"`
}

type PatientListFilter struct {
	IdentityNumber int
	Name           string
	PhoneNumber    string
	CreatedAt      string
	Limit          int
	Offset         int
}

type PatientListResult struct {
	IdentityNumber int    `json:"identityNumber"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phoneNumber"`
	CreatedAt      string `json:"createdAt"`
	Gender         string `json:"gender"`
	BirthDate      string `json:"birthDate"`
}

type PatientRegisterResponse struct {
	UserId string
}

type PatientRegisterResult struct {
	UserId string
}

type PatientRegisterParams struct {
	IdentityNumber int
	Name           string
	PhoneNumber    string
	BirthDate      string
	Gender         string
	IdentityCard   string
}

type UserListFilter struct {
	IdentityNumber int
	PhoneNumber    string
	Name           string
	CreatedAt      string
	Limit          int
	Offset         int
	UserId         string
}
