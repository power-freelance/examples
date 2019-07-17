package fserv

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

const (
	defaultMaxMemory = 32 << 20 // 32 MB
)

type (
	Result struct {
		Name     string `json:"name,omitempty"`
		Size     int64  `json:"size,omitempty"`
		MimeType string `json:"mime_type,omitempty"`
		Ext      string `json:"ext,omitempty"`
		Error    string `json:"error,omitempty"`
	}
)

func uploadMultipartFormData(c echo.Context, flags *Flags) error {

	if err := c.Request().ParseMultipartForm(defaultMaxMemory); err != nil {
		return err
	}

	resultSet := make(map[string][]*Result)

	for group, multipartFiles := range c.Request().MultipartForm.File {

		// Create sub dir if not exists
		subStaticRoot := path.Join(flags.StaticRoot, strings.Trim(group, "./"))

		if err := checkDir(subStaticRoot, 0755); err != nil {
			return err
		}

		// Prepare result set items
		resultSet[group] = make([]*Result, len(multipartFiles))

		// Parse each file in multipart data
		for i, multipartFile := range multipartFiles {
			resultSet[group][i] = processMultipartFile(flags, group, multipartFile)
		}
	}

	return c.JSON(http.StatusOK, resultSet)
}
func processMultipartFile(flags *Flags, group string, multipartFile *multipart.FileHeader) *Result {

	// Try open temporary uploaded file
	file, err := multipartFile.Open()
	if err != nil {
		return &Result{Error: err.Error()}
	}
	defer file.Close()

	// Read file header for mime detect
	head := make([]byte, 261)
	if bytes, err := file.Read(head); err != nil {
		return &Result{Error: err.Error()}
	} else if bytes != 261 {
		return &Result{Error: "error read file head"}
	}

	// Seek
	if newOffset, err := file.Seek(0, io.SeekStart); err != nil {
		return &Result{Error: err.Error()}
	} else if newOffset != 0 {
		return &Result{Error: "error seek to start file"}
	}

	// Detect mime type and extension
	mimeType, ext := mimetype.Detect(head)
	filename := strings.Join([]string{uuid.New().String(), ext}, ".")

	// Create destination file
	filepath := path.Join(flags.StaticRoot, group, filename)
	destFile, err := os.Create(filepath)
	if err != nil {
		return &Result{Error: err.Error()}
	}
	defer destFile.Close()

	// Copy to destination
	bytes, err := io.Copy(destFile, file)
	if err != nil {
		return &Result{Error: err.Error()}
	} else if bytes != multipartFile.Size {
		return &Result{Error: "error write destination file"}
	}

	return &Result{Name: path.Join(group, filename), Ext: ext, MimeType: mimeType, Size: multipartFile.Size}
}
