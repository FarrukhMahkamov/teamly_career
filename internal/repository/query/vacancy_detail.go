package query

const (
	// GetVacancyDetails is the query to get all vacancy details
	GetVacancyDetail = `SELECT * FROM tbl_vacancy_detail WHERE vacancy_id = ?`
)
