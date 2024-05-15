package user

type ITRegister struct {
	Nip      int64  `json:"nip" validate:"required"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}
