package entities

type NurseRegister struct {
	Nip           int64  `json:"nip" validate:"required"`
	Name          string `json:"name" validate:"required,min=5,max=50"`
	IdentityImage string `json:"identityCardScanImg" validate:"required"`
}

type NurseRegisterResponse struct {
	UserId string
}

type NurseRegisterParams struct {
	Nip          int64
	Name         string
	IdentityCard string
	Role         string
}

type NurseRegisterAccess struct {
	Password string `json:"password" validate:"required,min=5,max=33"`
	UserId   string `json:"userId"`
}
type NurseRegisterResult struct {
	UserId string `json:"userId"`
	Nip    int64  `json:"nip"`
	Name   string `json:"name"`
}

type NurseDeleteParams struct {
	UserId string
}

type NurseEditParams struct {
	Nip    int64  `json:"nip"`
	Name   string `json:"name"`
	UserId string `json:"userId"`
}

type NurseEdit struct {
	Nip  int64  `json:"nip" validate:"required"`
	Name string `json:"name" validate:"required,min=5,max=50"`
}
