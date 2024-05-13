package user

type ITRegister struct {
	Nip      string `json:"nip" binding:"required,min=5,max=13"`
	Name     string `json:"name" binding:"required,min=5,max=50"`
	Role     string `json:"role" binding:"required,min=2,max=5"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}
