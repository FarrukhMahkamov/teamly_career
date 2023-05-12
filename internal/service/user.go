package service

import (
	"errors"

	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository"
	"github.com/FarrukhMahkamov/teamly_career/pkg"
)

type UserService struct {
	repository repository.User
}

// NewUserService ...
func NewUserService(repository repository.User) *UserService {
	return &UserService{repository: repository}
}

// RegistrUser ...
func (s *UserService) RegisterUser(UserRequest core.UserRequest) (int64, error) {
	//Hash password
	UserPassword, err := pkg.HashPassword(UserRequest.UserPassword)
	if err != nil {
		return 0, err
	}

	//Set hash password
	UserRequest.UserPassword = UserPassword

	//Execute query
	return s.repository.RegistrUser(UserRequest)
}

// UserPersonalInfoUpdate ...
func (s *UserService) UserPersonalInfoUpdate(UserPersonalInfoUpdateRequest core.UserPersonalInfoUpdateRequest, UserID int64) error {
	return s.repository.UserPersonalInfoUpdate(UserPersonalInfoUpdateRequest, UserID)
}

// UpdateUserPassword ...
func (s *UserService) UpdateUserPassword(UserID int64, UpdateUserPasswordRequest core.UpdateUserPasswordRequest) error {
	UserPassword, err := s.repository.GetUserPassword(UserID)
	if err != nil {
		return err
	}

	if err := pkg.ComparePassword(UserPassword, UpdateUserPasswordRequest.OldPassword); err != nil {
		return errors.New("old password is incorrect")
	}

	NewUserPassword, err := pkg.HashPassword(UpdateUserPasswordRequest.NewPassword)
	if err != nil {
		return err
	}

	if err := s.repository.UpdateUserPassword(UserID, NewUserPassword); err != nil {
		return err
	}

	return nil
}
