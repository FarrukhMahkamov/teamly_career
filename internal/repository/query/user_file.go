package query

const (
	AddUserFile = `INSERT INTO tbl_user_file 
		(user_id, file_name, file_size, file_type, status_id)
		VALUES
		($1, $2, $3, $4, $5)`
)
