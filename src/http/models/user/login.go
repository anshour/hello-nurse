package user

type UserLogin struct {
	Nip      int64  `json:"nip" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}
