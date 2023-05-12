package pkg

import (
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

	DestinationFile, err := os.Create(filepath.Join(Destination, NewFileName))
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
