package v1routes

import (
	uploadController "hello-nurse/src/http/controllers/upload"
	middleware "hello-nurse/src/utils/middlewares"
)

func (i *V1Routes) MountUpload() {

	upload := uploadController.New(&uploadController.V1Upload{
		DB: i.DB,
	})
	i.Echo.POST("/image", upload.UploadImage, middleware.Authentication())

}
