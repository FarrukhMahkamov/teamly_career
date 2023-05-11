package core

import "time"

type Vacancy struct {
	VacancyId          int        `json:"vacancy_id" db:"vacancy_id"`
	VacancyTeamId      int        `json:"vacancy_team_id" db:"vacancy_team_id"`
	VacancyStatusId    int        `json:"vacancy_status_id" db:"vacancy_status_id"`
	Position           string     `json:"position" db:"position"`
	VacancyDescription string     `json:"vacancy_description" db:"vacancy_description"`
	CreatedAt          time.Time  `json:"created_at" db:"created_at" format:"15:04:05"`
	UpdatedAt          time.Time  `json:"updated_at" db:"updated_at" format:"15:04:05"`
	DeletedAt          *time.Time `json:"deleted_at" db:"deleted_at"`
}

type VacancyRequest struct {
	VacancyTeamId      int    `json:"vacancy_team_id" db:"vacancy_team_id"`
	VacancyStatusId    int    `json:"vacancy_status_id" db:"vacancy_status_id"`
	Position           string `json:"position" db:"position"`
	VacancyDescription string `json:"vacancy_description" db:"vacancy_description"`
}

type AddVacancyRequest struct {
	VacancyTeamId      int    `json:"vacancy_team_id" db:"vacancy_team_id"`
	VacancyStatusId    int    `json:"vacancy_status_id" db:"vacancy_status_id"`
	Position           string `json:"position" db:"position"`
	VacancyDescription string `json:"vacancy_description" db:"vacancy_description"`
	Level              string `json:"level" db:"level"`
	Experience         string `json:"experience" db:"experience"`
	WorkType           string `json:"work_type" db:"work_type"`
	WorkTime           string `json:"work_time" db:"work_time"`
	WorkLocation       string `json:"work_location" db:"work_location"`
	Salary             string `json:"salary" db:"salary"`
}
