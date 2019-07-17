package fserv

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func Upload(flags *Flags) echo.HandlerFunc {
	return func(c echo.Context) error {

		contentType := c.Request().Header.Get("Content-Type")

		if strings.HasPrefix(contentType, "multipart/form-data") {
			return uploadMultipartFormData(c, flags)
		} else if contentType == "application/octet-stream" {
			return uploadOctetStream(c, flags)
		} else {
			return c.NoContent(http.StatusBadRequest)
		}
	}
}
