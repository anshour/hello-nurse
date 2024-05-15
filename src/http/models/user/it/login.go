package user

type ITLogin struct {
	Nip      int64  `json:"nip" binding:"required"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}
