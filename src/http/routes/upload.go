package v1routes

import (
	uploadController "hello-nurse/src/http/controllers/upload"
)

func (i *V1Routes) MountUpload() {
	g := i.Echo.Group("")

	upload := uploadController.New(&uploadController.V1Upload{
		DB: i.DB,
	})
	g.POST("/image", upload.UploadImage)

}
