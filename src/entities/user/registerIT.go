package entities

type ITRegister struct {
	Nip      int64  `json:"nip" validate:"required"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type ITRegisterResponse struct {
	Id string
}

type ITRegisterParams struct {
	Nip      int64
	Name     string
	Password string
	Role     string
}

type ITRegisterResult struct {
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
	Nip         int64  `json:"nip"`
	Name        string `json:"name"`
}
