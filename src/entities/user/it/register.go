package entities

type ITRegisterResponse struct {
	Id string
}

type ParamsITRegister struct {
	Nip      int64
	Name     string
	Password string
	Role     string
}

type ResultITRegister struct {
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
	Nip         int64  `json:"nip"`
	Name        string `json:"name"`
}
