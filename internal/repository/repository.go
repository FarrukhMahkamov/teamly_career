package repository

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/jmoiron/sqlx"
)

type Team interface {
	GetTeams() ([]core.Team, error)
}

type Auth interface {
	LoginUser(UserLoginRequest core.UserLoginRequest) (*core.User, error)
}

type User interface {
	RegistrUser(UserRequest core.UserRequest) (int64, error)
	UserPersonalInfoUpdate(UserPersonalInfoUpdateRequest core.UserPersonalInfoUpdateRequest, UserID int64) error
	GetUserPassword(UserID int64) (string, error)
	UpdateUserPassword(UserID int64, UserPassword string) error
}

type Vacancy interface {
	GetVacancies() ([]core.Vacancy, error)
	AddVacancy(VacancyRequest core.VacancyRequest) (int64, error)
}

type VacancyDetail interface {
	GetVacancyDetails(VacancyID int64) (*core.VacancyDetail, error)
	AddVacancyDetail(VacancyDetailRequest core.VacancyDetailRequest) (int64, error)
}

type UserFile interface {
	AddUserFile(UserFileRequest core.UserFileRequest) error
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
		User:          NewUserRepository(db),
		Auth:          NewAuthRepository(db),
		UserFile:      NewUserFileRepository(db),
	}
}
