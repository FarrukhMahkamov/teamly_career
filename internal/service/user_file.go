package service

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository"
)

type UserFileService struct {
	repository repository.UserFile
}

func NewUserFileService(repository *repository.Repository) *UserFileService {
	return &UserFileService{repository: repository.UserFile}
}

// AddUserFile adds user file to database
func (s *UserFileService) AddUserFile(UserFileRequest core.UserFileRequest) error {
	return s.repository.AddUserFile(UserFileRequest)
}
