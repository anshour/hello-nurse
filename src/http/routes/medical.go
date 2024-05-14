package v1routes

import (
	medicalController "hello-nurse/src/http/controllers/medical"
)

func (i *V1Routes) MountMedical() {
	g := i.Echo.Group("/medical")

	medicalController := medicalController.New(&medicalController.V1Medical{
		DB: i.DB,
	})

	g.GET("/patient", medicalController.PatientList)
	g.POST("/patient", medicalController.PatientRegister)
	g.GET("/record", medicalController.RecordList)
	g.POST("/record", medicalController.RecordRegister)

}
