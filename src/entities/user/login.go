package entities

type LoginResponse struct {
	Id       string
	Nip      int64
	Name     string
	Password string
	Role     string
}

type LoginParams struct {
	Nip      int64
	Password string
}

type LoginResult struct {
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Nip         int64  `json:"nip"`
}
