package entities

type NurseRegisterResponse struct {
	UserId string
}

type NurseRegisterParams struct {
	Nip          int64
	Name         string
	IdentityCard string
	Role         string
}

type NurseEditParams struct {
	Nip  int64
	Name string
}
type NurseRegisterAccess struct {
	Password string
	UserId   string
}
type NurseRegisterResult struct {
	UserId string `json:"userId"`
	Nip    int64  `json:"nip"`
	Name   string `json:"name"`
}
