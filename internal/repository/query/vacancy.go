package query

const (
	// GetVacancies returns all vacancies
	GetVacancies = `SELECT * FROM tbl_vacancy WHERE deleted_at IS NULL`

	// AddVacancy adds a new vacancy
	AddVacancy = `INSERT INTO tbl_vacancy 
	(vacancy_team_id, vacancy_status_id, position, vacancy_description) 
	VALUES ($1, $2, $3, $4) RETURNING vacancy_id`
)
