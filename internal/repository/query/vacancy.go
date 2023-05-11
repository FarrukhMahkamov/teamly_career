package query

const (
	// GetVacancies returns all vacancies
	GetVacancies = `SELECT * FROM tbl_vacancy WHERE deleted_at IS NULL`
)
