package entities

type ITLoginResponse struct {
	Id       string
	Nip      int64
	Name     string
	Password string
	Role     string
}

type ITLoginParams struct {
	Nip      int64
	Password string
}

type ITLoginResult struct {
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
	Name        string `json:"name"`
}
