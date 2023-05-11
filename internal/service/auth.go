package service

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository"
	"github.com/FarrukhMahkamov/teamly_career/pkg"
)

type AuthService struct {
	repository repository.Auth
}

func NewAuthService(repository repository.Auth) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) LoginUser(UserLoginRequest core.UserLoginRequest) (*core.User, error) {
	//Get user by email
	User, err := s.repository.LoginUser(UserLoginRequest)
	if err != nil {
		return nil, err
	}

	//Check password
	err = pkg.CheckPassword(UserLoginRequest.UserPassword, User.UserPassword)
	if err != nil {
		return nil, err
	}

	//Generate token
	Token, err := pkg.GenerateToken(User.UserID)
	if err != nil {
		return nil, err
	}

	//Set token
	User.Token = Token

	//Return user
	return User, nil
}
