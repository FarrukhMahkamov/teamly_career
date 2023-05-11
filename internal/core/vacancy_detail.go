package core

import "time"

type VacancyDetail struct {
	VacancyDetailId int       `json:"vacancy_detail_id" db:"vacancy_detail_id"`
	VacancyId       int       `json:"vacancy_id" db:"vacancy_id"`
	ApplyCount      int       `json:"apply_count" db:"apply_count"`
	Level           string    `json:"level" db:"level"`
	Experience      string    `json:"experience" db:"experience"`
	WorkType        string    `json:"work_type" db:"work_type"`
	WorkTime        string    `json:"work_time" db:"work_time"`
	WorkLocation    string    `json:"work_location" db:"work_location"`
	CreatedAt       time.Time `json:"created_at" db:"created_at" format:"15:04:05"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at" format:"15:04:05"`
}
