package entities

type RecordRegisterResponse struct {
	UserId string
}

type RecordRegisterResult struct {
	UserId string
}

type RecordRegisterParams struct {
	IdentityNumber int64
	Symptoms       string
	Medications    string
}
