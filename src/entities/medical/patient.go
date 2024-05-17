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

type UserListFilter struct {
	IdentityNumber int
	PhoneNumber    string
	Name           string
	CreatedAt      string
	Limit          int
	Offset         int
	UserId         string
}
