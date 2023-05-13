package repository

import (
	"context"
	"time"

	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type UserFileRepository struct {
	db *sqlx.DB
}

func NewUserFileRepository(db *sqlx.DB) *UserFileRepository {
	return &UserFileRepository{db: db}
}

// Add user file
func (r *UserFileRepository) AddUserFile(UserFileRequest core.UserFileRequest) error {
	// Set timeout for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Execute query
	_, err := r.db.ExecContext(ctx, query.AddUserFile,
		UserFileRequest.UserID,
		UserFileRequest.FileName,
		UserFileRequest.FileSize,
		UserFileRequest.FileType,
		UserFileRequest.StatusID,
	)

	// Check for errors
	if err != nil {
		return err
	}

	// Return nil error
	return nil
}
