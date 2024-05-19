package uploadController

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (dbase *V1Upload) UploadImage(c echo.Context) (err error) {
	// Read form file
	file, err := c.FormFile("file")
	println(err.Error())
	if err != nil {
		return err
	}
	// Source
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Success",
		Data:    "",
	})

}
