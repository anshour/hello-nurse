package entities

type UserListResponse struct {
	Id        string `json:"userId"`
	Nip       int    `json:"nip"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
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
