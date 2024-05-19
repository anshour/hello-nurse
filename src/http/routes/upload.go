package v1routes

import (
	uploadController "hello-nurse/src/http/controllers/upload"
	middleware "hello-nurse/src/utils/middlewares"
)

func (i *V1Routes) MountUpload() {
	g := i.Echo.Group("")

	upload := uploadController.New(&uploadController.V1Upload{
		DB: i.DB,
	})
	g.Use(middleware.Authentication())
	g.POST("/image", upload.UploadImage)

}
