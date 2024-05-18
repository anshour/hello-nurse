package entities

type RecordRegister struct {
	IdentityNumber int    `json:"identityNumber" validate:"required"`
	Symptoms       string `json:"symptoms" validate:"required,min=5,max=2000"`
	Medications    string `json:"medications" validate:"required,min=5,max=2000"`
}

type RecordRegisterResponse struct {
	Id string `json:"recordId"`
}

type RecordRegisterParams struct {
	UserId         string
	IdentityNumber int
	Symptoms       string
	Medications    string
}

type RecordListFilter struct {
	IdentityNumber int
	UserId         string
	Nip            string
	CreatedAt      string
	Limit          int
	Offset         int
}

type IdentityDetail struct {
	IdentityNumber int    `json:"identityNumber"`
	PhoneNumber    string `json:"phoneNumber"`
	Name           string `json:"name"`
	BirthDate      string `json:"birthDate"`
	Gender         string `json:"gender"`
	IdentityImage  string `json:"identityCardScanImg"`
}

type CreatedBy struct {
	Nip    int    `json:"nip"`
	Name   string `json:"name"`
	UserId string `json:"userId"`
}

type RecordListResult struct {
	IdentityDetail IdentityDetail `json:"identityDetail"`
	Symptoms       string         `json:"symptoms"`
	Medications    string         `json:"medications"`
	CreatedAt      string         `json:"createdAt"`
	CreatedBy      CreatedBy      `json:"createdBy"`
}
