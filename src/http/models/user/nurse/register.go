package user

type NurseRegister struct {
	Nip                 int64  `json:"nip" binding:"required"`
	Name                string `json:"name" binding:"required,min=5,max=50"`
	IdentityCardScanImg string `json:"identity_image" binding:"required"`
}
