package entities

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
