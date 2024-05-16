package entities

type NurseRegisterResponse struct {
	Id string
}

type NurseRegisterParams struct {
	Nip          int64
	Name         string
	IdentityCard string
}

type NurseRegisterResult struct {
	UserId string `json:"userId"`
	Nip    int64  `json:"nip"`
	Name   string `json:"name"`
}
