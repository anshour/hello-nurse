package user

type NurseRegister struct {
	Nip           int64  `json:"nip" validate:"required"`
	Name          string `json:"name" validate:"required,min=5,max=50"`
	IdentityImage string `json:"identityCardScanImg" validate:"required"`
}
