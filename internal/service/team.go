package service

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository"
)

type TeamService struct {
	repository *repository.Repository
}

func NewTeamService(repository *repository.Repository) *TeamService {
	return &TeamService{repository: repository}
}

func (s *TeamService) GetTeams() ([]core.Team, error) {
	return s.repository.Team.GetTeams()
}
