package repository

import (
	"context"
	"time"

	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type VacancyRepository struct {
	db *sqlx.DB
}

func NewVacancyRepository(db *sqlx.DB) *VacancyRepository {
	return &VacancyRepository{db: db}
}

// GetVacancies returns all vacancies
func (r *VacancyRepository) GetVacancies() ([]core.Vacancy, error) {
	// Set timeout for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a slice of Vacancy to hold the query results.
	var Vacancies []core.Vacancy

	// Execute query
	rows, err := r.db.QueryxContext(ctx, query.GetVacancies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		// Declare a Vacancy to hold each row's data
		var Vacancy core.Vacancy

		// Unmarshal the row's data into vacancy
		if err := rows.StructScan(&Vacancy); err != nil {
			return nil, err
		}

		// Append vacancy to vacancies
		Vacancies = append(Vacancies, Vacancy)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return result and nil error
	return Vacancies, nil
}
