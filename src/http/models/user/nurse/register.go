package user

type NurseRegister struct {
	Nip                 string `json:"nip" binding:"required,min=5,max=13"`
	Name                string `json:"name" binding:"required,min=5,max=50"`
	IdentityCardScanImg string `json:"role" binding:"required,min=2,max=5"`
}
