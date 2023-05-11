package core

type Team struct {
	TeamID   int64  `json:"team_id" db:"team_id"`
	TeamName string `json:"team_name" db:"team_name"`
}
