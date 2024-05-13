package medical

type RecordRegister struct {
	IdentityNumber string `json:"identity_number" binding:"required,min=13,max=13"`
	Symptoms       string `json:"symptoms" binding:"required,min=5,max=255"`
	Medications    string `json:"medications" binding:"required,min=5,max=255"`
}
