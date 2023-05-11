package service

import (
	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository"
)

type VacancyDetailService struct {
	repository repository.VacancyDetail
}

func NewVacancyDetailService(repository repository.VacancyDetail) *VacancyDetailService {
	return &VacancyDetailService{repository: repository}
}

func (s *VacancyDetailService) GetVacancyDetails(VacancyID int64) (*core.VacancyDetail, error) {
	return s.repository.GetVacancyDetails(VacancyID)
}

func (s *VacancyDetailService) AddVacancyDetail(VacancyDetailRequest core.VacancyDetailRequest) (int64, error) {
	return s.repository.AddVacancyDetail(VacancyDetailRequest)
}
