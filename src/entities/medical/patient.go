package entities

type PatientRegisterResponse struct {
	UserId string
}

type PatientRegisterResult struct {
	UserId string
}

type PatientRegisterParams struct {
	IdentityNumber string
	Name           string
	PhoneNumber    string
	BirthDate      string
	Gender         string
	IdentityCard   string
}
