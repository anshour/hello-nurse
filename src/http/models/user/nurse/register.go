package user

type NurseRegister struct {
	Nip                 int64  `json:"nip" validate:"required"`
	Name                string `json:"name" validate:"required,min=5,max=50"`
	IdentityCardScanImg string `json:"identity_image" validate:"required"`
}
