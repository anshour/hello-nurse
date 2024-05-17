package v1routes

import (
	medicalController "hello-nurse/src/http/controllers/medical"
	middleware "hello-nurse/src/utils/middlewares"
)

func (i *V1Routes) MountUpload() {
	g := i.Echo.Group("/medical")

	medicalController := medicalController.New(&medicalController.V1Medical{
		DB: i.DB,
	})
	g.Use(middleware.Authentication())
	g.POST("/image", medicalController.PatientList)

}
