package repository

import (
	"context"
	"time"

	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// RegistrUser ...
func (r *UserRepository) RegistrUser(UserRequest core.UserRequest) (int64, error) {
	//set timout for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Declare var for hold user_id
	var UserID int64

	//Execute query and scan user_id
	err := r.db.QueryRowContext(ctx, query.RegistrUser,
		UserRequest.UserName,
		UserRequest.UserSecondName,
		UserRequest.UserEmail,
		UserRequest.UserPhone,
		UserRequest.UserPassword,
		UserRequest.UserPhoto,
	).Scan(&UserID)

	//Check error
	if err != nil {
		return 0, err
	}

	//Return user_id
	return UserID, nil
}

// UserPersonalInfoUpdate ...
func (r *UserRepository) UserPersonalInfoUpdate(UserPersonalInfoUpdateRequest core.UserPersonalInfoUpdateRequest, UserID int64) error {
	//set timout for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Execute query
	_, err := r.db.ExecContext(ctx, query.UserPersonalInfoUpdate,
		UserPersonalInfoUpdateRequest.UserName,
		UserPersonalInfoUpdateRequest.UserSecondName,
		UserPersonalInfoUpdateRequest.UserEmail,
		UserPersonalInfoUpdateRequest.UserPhone,
		UserID,
	)

	//Check error
	if err != nil {
		return err
	}

	//Return nil
	return nil
}

// GetUserPassword ...
func (r *UserRepository) GetUserPassword(UserID int64) (string, error) {
	//set timout for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Declare var for hold user_password
	var UserPassword string

	//Execute query and scan user_password
	err := r.db.QueryRowContext(ctx, query.GetUserPassword, UserID).Scan(&UserPassword)

	//Check error
	if err != nil {
		return "", err
	}

	//Return user_password
	return UserPassword, nil
}

// UpdateUserPassword ...
func (r *UserRepository) UpdateUserPassword(UserID int64, UserPassword string) error {
	//set timout for query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Execute query
	_, err := r.db.ExecContext(ctx, query.UpdateUserPassword, UserPassword, UserID)

	//Check error
	if err != nil {
		return err
	}

	//Return nil
	return nil
}
