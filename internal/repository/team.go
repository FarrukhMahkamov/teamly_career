package repository

import (
	"context"
	"time"

	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type TeamRepository struct {
	db *sqlx.DB
}

func NewTeamRepository(db *sqlx.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

// GetTeams returns all teams
func (r *TeamRepository) GetTeams() ([]core.Team, error) {
	// Set timeout for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a slice of Team to hold the query results.
	var teams []core.Team

	// Execute query
	rows, err := r.db.QueryxContext(ctx, query.GetTeams)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		var team core.Team
		err := rows.StructScan(&team)
		if err != nil {
			return nil, err
		}

		// Append team to slice
		teams = append(teams, team)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return result and nil error
	return teams, nil
}
