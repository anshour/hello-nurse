package entities

type UserListResponse struct {
	Id        string
	Nip       int
	Name      string
	Role      string
	CreatedAt string
}

type UserListFilter struct {
	Id        string
	Nip       int
	Name      string
	Role      string
	CreatedAt string
	Limit     int
	Offset    int
}

type UserListResult struct {
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
	Name        string `json:"name"`
}
