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
	GetVacancies() ([]core.Vacancy, error)
}

type VacancyDetail interface {
	GetVacancyDetails(VacancyID int64) (*core.VacancyDetail, error)
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
		Team:          NewTeamService(repository),
		Vacancy:       NewVacancyService(repository),
		VacancyDetail: NewVacancyDetailService(repository),
	}
}
