package user

type NurseEdit struct {
	Nip  int64  `json:"nip" validate:"required"`
	Name string `json:"name" validate:"required,min=5,max=50"`
}
