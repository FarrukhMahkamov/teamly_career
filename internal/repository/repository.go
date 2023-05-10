package repository

import "github.com/jmoiron/sqlx"

type TeamRepository interface {
}

type JobCategoryRepository interface {
}

type AuthRepository interface {
}

type UserRepository interface {
}

type VacancyRepository interface {
}

type VacancyDetailRepository interface {
}

type UserFileRepository interface {
}

type Repository struct {
	TeamRepository
	JobCategoryRepository
	AuthRepository
	UserRepository
	VacancyRepository
	VacancyDetailRepository
	UserFileRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
