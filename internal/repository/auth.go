package repository

import (
	"context"
	"time"

	"github.com/FarrukhMahkamov/teamly_career/internal/core"
	"github.com/FarrukhMahkamov/teamly_career/internal/repository/query"
	"github.com/jmoiron/sqlx"
)

type AuthRepositort struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepositort {
	return &AuthRepositort{db: db}
}

// LoginUser ...
func (r *AuthRepositort) LoginUser(UserLoginRequest core.UserLoginRequest) (*core.User, error) {
	//Set timout for query
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//Declare struct for hold user data
	User := &core.User{}

	//Execute query and scan user data
	err := r.db.QueryRowContext(ctx, query.LoginUser, UserLoginRequest.UserEmail).Scan(
		&User.UserID,
		&User.UserName,
		&User.UserSecondName,
		&User.UserEmail,
		&User.UserPhone,
		&User.UserPassword,
		&User.UserPhoto,
	)

	//Check error
	if err != nil {
		return nil, err
	}

	//Return user_id
	return User, nil
}
