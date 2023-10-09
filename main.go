package gouploader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type UploadedFile struct {
	Filename string `json:"filename"`
}

func Single(r *http.Request, directory string, fieldName string) (*UploadedFile, error) {
	err := r.ParseMultipartForm(30)
	if err != nil {
		return nil, err
	}

	form := r.MultipartForm

	if len(form.File[fieldName]) == 0 {
		return nil, nil
	}

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	headers := form.File[fieldName][0]
	file, err := headers.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	destFile, err := os.Create(filepath.Join(directory, headers.Filename))
	if err != nil {
		return nil, err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return nil, err
	}

	uploadedFile := &UploadedFile{
		Filename: fmt.Sprintf("%s/%s", directory, headers.Filename),
	}

	return uploadedFile, nil
}

func Multiple(r *http.Request, directory string, fieldName string) ([]UploadedFile, error) {
	err := r.ParseMultipartForm(30)
	if err != nil {
		return nil, err
	}

	form := r.MultipartForm

	var uploadedFiles []UploadedFile

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	for _, headers := range form.File[fieldName] {
		file, err := headers.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		destFile, err := os.Create(filepath.Join(directory, headers.Filename))
		if err != nil {
			return nil, err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, file)
		if err != nil {
			return nil, err
		}

		uploadedFiles = append(uploadedFiles, UploadedFile{
			Filename: fmt.Sprintf("%s/%s", directory, headers.Filename),
		})
	}

	return uploadedFiles, nil
}
