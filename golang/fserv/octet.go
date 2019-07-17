package fserv

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func uploadOctetStream(c echo.Context, flags *Flags) error {

	// Generate temporary file
	tmpFile, err := ioutil.TempFile("", "upload")
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	// Copy body in temp file
	if bytes, err := io.Copy(tmpFile, c.Request().Body); err != nil {
		return err
	} else if bytes != c.Request().ContentLength {
		return errors.New("error copy temp file")
	}

	// Read file header for mime detect
	head := make([]byte, 261)
	if bytes, err := tmpFile.ReadAt(head, 0); err != nil {
		return err
	} else if bytes != 261 {
		return errors.New("error read file head")
	}

	// Detect mime type and extension
	mimeType, ext := mimetype.Detect(head)
	filename := strings.Join([]string{uuid.New().String(), ext}, ".")

	// Create destination file
	filepath := path.Join(flags.StaticRoot, filename)
	destFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Seek
	if newOffset, err := tmpFile.Seek(0, io.SeekStart); err != nil {
		return err
	} else if newOffset != 0 {
		return errors.New("error seek to start file")
	}

	// Copy to destination
	bytes, err := io.Copy(destFile, tmpFile)
	if err != nil {
		return err
	} else if bytes != c.Request().ContentLength {
		return errors.New("copy destination error")
	}

	return c.JSON(http.StatusOK, &Result{Name: filename, Ext: ext, MimeType: mimeType, Size: c.Request().ContentLength})
}
