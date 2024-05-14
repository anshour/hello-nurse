package user

type ITRegister struct {
	Nip      string `json:"nip" binding:"required,min=5,max=13"`
	Name     string `json:"name" binding:"required,min=5,max=50"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}
