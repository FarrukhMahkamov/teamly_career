package query

const (
	RegistrUser = `
		INSERT INTO tbl_user 
		(user_name, user_second_name, user_email, user_phone, user_password, user_photo)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id`

	UserPersonalInfoUpdate = `UPDATE tbl_user SET user_name=$1, user_second_name=$2, user_email=$3, user_phone=$4 WHERE user_id=$5`
	GetUserPassword        = `SELECT user_password FROM tbl_user WHERE user_id=$1`
	UpdateUserPassword     = `UPDATE tbl_user SET user_password=$1 WHERE user_id=$2`
)
