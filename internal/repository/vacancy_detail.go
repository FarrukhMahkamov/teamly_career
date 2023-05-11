package repository

import (
	"context"
	"time"

	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type VacancyDetailRepository struct {
	db *sqlx.DB
}

func NewVacancyDetailRepository(db *sqlx.DB) *VacancyDetailRepository {
	return &VacancyDetailRepository{db: db}
}

func (r *VacancyDetailRepository) GetVacancyDetails(VacancyID int64) (*core.VacancyDetail, error) {
	//Set timeout for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Declare a slice of VacancyDetail to hold the query results.
	VacancyDetail := &core.VacancyDetail{}

	//Execute query
	err := r.db.QueryRowContext(ctx, query.GetVacancyDetail, VacancyID).Scan(
		&VacancyDetail.VacancyDetailId,
		&VacancyDetail.VacancyId,
		&VacancyDetail.ApplyCount,
		&VacancyDetail.Level,
		&VacancyDetail.Experience,
		&VacancyDetail.WorkType,
		&VacancyDetail.WorkTime,
		&VacancyDetail.WorkLocation,
		&VacancyDetail.CreatedAt,
		&VacancyDetail.UpdatedAt,
	)

	//Check for errors during row scan
	if err != nil {
		return nil, err
	}

	//Return result and nil error
	return VacancyDetail, nil
}

func (r *VacancyDetailRepository) AddVacancyDetail(VacancyDetailRequest core.VacancyDetailRequest) (int64, error) {
	//Set timeout for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Declare a VacancyDetailID to hold the value of the last inserted id
	var VacancyDetailID int64

	//Execute query and return VacancyDetailID
	err := r.db.QueryRowContext(ctx, query.AddVacancyDetail,
		VacancyDetailRequest.VacancyId,
		VacancyDetailRequest.Level,
		VacancyDetailRequest.Experience,
		VacancyDetailRequest.WorkType,
		VacancyDetailRequest.WorkTime,
		VacancyDetailRequest.WorkLocation,
		VacancyDetailRequest.Salary,
	).Scan(&VacancyDetailID)

	//Check for errors during row scan
	if err != nil {
		return 0, err
	}

	//Return result and nil error
	return VacancyDetailID, nil
}
