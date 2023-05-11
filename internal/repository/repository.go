package repository

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	Team
	Auth
	User
	Vacancy
	VacancyDetail
	UserFile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Team:          NewTeamRepository(db),
		Vacancy:       NewVacancyRepository(db),
		VacancyDetail: NewVacancyDetailRepository(db),
	}
}
