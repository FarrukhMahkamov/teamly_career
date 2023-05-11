package service

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository"
)

type Team interface {
	GetTeams() ([]core.Team, error)
}

type Auth interface {
}

type User interface {
}

type Vacancy interface {
}

type VacancyDetail interface {
}

type UserFile interface {
}

type Service struct {
	Team
	Auth
	User
	Vacancy
	VacancyDetail
	UserFile
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Team: NewTeamService(repository),
	}
}
