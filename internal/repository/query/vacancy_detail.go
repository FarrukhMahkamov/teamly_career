package query

const (
	// GetVacancyDetails is the query to get all vacancy details
	GetVacancyDetail = `SELECT * FROM tbl_vacancy_detail WHERE vacancy_id = ?`

	//AddVacancyDetail is the query to add a new vacancy detail
	AddVacancyDetail = `INSERT INTO tbl_vacancy_detail
		(vacancy_id, level, experience, work_type, work_time, work_location, salary)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING vacancy_detail_id`
)
