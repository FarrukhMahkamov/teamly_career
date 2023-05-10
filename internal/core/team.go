package core

import "time"

type Team struct {
	TeamID    int64     `json:"team_id"`
	TeamName  string    `json:"team_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
