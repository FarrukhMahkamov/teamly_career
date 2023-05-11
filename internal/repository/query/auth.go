package query

const LoginUser = `SELECT user_id, user_name, user_second_name, user_email, user_phone, user_password, user_photo FROM tbl_user WHERE user_email = $1`
