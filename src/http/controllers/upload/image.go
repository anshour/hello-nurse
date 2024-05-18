package uploadController

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (dbase *V1Upload) UploadImage(c echo.Context) (err error) {

	form, err := c.MultipartForm()
	if err != nil {
		println("error - 1", err)
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			println("error - 2", err)

			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(file.Filename)
		if err != nil {
			println("error - 3", err)

			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}
	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Success",
		Data:    "",
	})

}
