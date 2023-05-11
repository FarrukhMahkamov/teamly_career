package service

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository"
)

type VacancyService struct {
	repository repository.Vacancy
}

func NewVacancyService(repository *repository.Repository) *VacancyService {
	return &VacancyService{repository: repository.Vacancy}
}

func (s *VacancyService) GetVacancies() ([]core.Vacancy, error) {
	return s.repository.GetVacancies()
}
