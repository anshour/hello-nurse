package user

type ITLogin struct {
	Nip      string `json:"nip" binding:"required,min=5,max=13"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}
