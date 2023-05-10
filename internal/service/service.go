package service

import "github.com/FarrukhMahkamov/teamly_career/internal/repository"

type TeamService interface {
}

type JobCategory interface {
}

type AuthService interface {
}

type UserService interface {
}

type VacancyService interface {
}

type VacancyDetailService interface {
}

type UserFileService interface {
}

type Service struct {
	TeamService
	JobCategory
	AuthService
	UserService
	VacancyService
	VacancyDetailService
	UserFileService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
