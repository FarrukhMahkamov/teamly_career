package core

import "time"

type User struct {
	UserID         int64     `json:"user_id" db:"user_id"`
	UserName       string    `json:"user_name" db:"user_name"`
	UserSecondName string    `json:"user_second_name" db:"user_second_name"`
	UserEmail      string    `json:"user_email" db:"user_email"`
	UserPhone      string    `json:"user_phone" db:"user_phone"`
	UserPassword   string    `json:"user_password" db:"user_password"`
	UserPhoto      *string   `json:"user_photo" db:"user_photo"`
	UserStatus     int8      `json:"user_status" db:"user_status"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at" db:"deleted_at"`
}

type UserRequest struct {
	UserName       string `json:"user_name" db:"user_name" binding:"required" validate:"required,min=2,max=100"`
	UserSecondName string `json:"user_second_name" db:"user_second_name" binding:"required" validate:"required,min=2,max=100"`
	UserEmail      string `json:"user_email" db:"user_email" binding:"required" validate:"required,email"`
	UserPhone      string `json:"user_phone" db:"user_phone" binding:"required" validate:"required,min=9,max=13"`
	UserPassword   string `json:"user_password" db:"user_password" binding:"required" validate:"required,min=8,max=100"`
	UserPhoto      string `json:"user_photo" db:"user_photo" validate:"required,min=2,max=100"`
}

type UserLoginRequest struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type UserLoginResponse struct {
	UserID         int64     `json:"user_id"`
	UserName       string    `json:"user_name"`
	UserSecondName string    `json:"user_second_name"`
	UserEmail      string    `json:"user_email"`
	UserPhone      string    `json:"user_phone"`
	UserPhoto      string    `json:"user_photo"`
	UserStatus     int8      `json:"user_status"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
