package pkg

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func UploadRequestFile(File *multipart.FileHeader, Destination string) (string, error) {
	SourceFile, err := File.Open()
	if err != nil {
		return "", err
	}

	defer SourceFile.Close()

	NewFileName := uuid.New().String() + filepath.Ext(File.Filename)

	var Folder string

	switch Destination {
	case "cv":
		Folder = "assets/cv"
	case "photo":
		Folder = "assets/photo"
	default:
		return "", errors.New("'nvalid destination")
	}

	err = os.MkdirAll(Destination, os.ModePerm)
	if err != nil {
		return "", err
	}

	DestinationFile, err := os.Create(filepath.Join(Folder, NewFileName))
	if err != nil {
		return "", err
	}

	defer DestinationFile.Close()

	_, err = io.Copy(DestinationFile, SourceFile)
	if err != nil {
		return "", err
	}

	return NewFileName, nil
}
