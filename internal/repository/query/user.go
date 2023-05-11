package query

const (
	RegistrUser = `
		INSERT INTO tbl_user 
		(user_name, user_second_name, user_email, user_phone, user_password, user_photo)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id`
)
