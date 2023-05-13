package core

type UserFileRequest struct {
	UserID   int    `json:"user_id" db:"user_id"`
	FileName string `json:"file_names" db:"file_name"`
	FileSize int64  `json:"file_size" db:"file_size"`
	FileType string `json:"file_type" db:"file_type"`
	StatusID int    `json:"status_id" db:"status_id"`
}

type UserFile struct {
	UserFileId int    `json:"user_file_id" db:"user_file_id"`
	UserId     int    `json:"user_id" db:"user_id"`
	FileName   string `json:"file_names" db:"file_name"`
	FileSize   int64  `json:"file_size" db:"file_size"`
	FileType   string `json:"file_type" db:"file_type"`
	StatusID   int    `json:"status_id" db:"status_id"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
	DeletedAt  string `json:"deleted_at" db:"deleted_at"`
}
