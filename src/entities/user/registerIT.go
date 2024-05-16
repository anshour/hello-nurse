package entities

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
