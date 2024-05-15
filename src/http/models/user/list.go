package user

type UserList struct {
	UserId    string `json:"userId"`
	Nip       int    `json:"nip"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}
