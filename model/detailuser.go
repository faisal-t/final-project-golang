package model

type (
	DetailUser struct {
		ID       int    `json:"id"`
		User_Id  int    `json:"user_id"`
		FullName string `json:"full_name"`
		Alamat   string `json:"alamat"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
		// CreatedAt time.Time `json:"created_at"`
		// UpdatedAt time.Time `json:"updated_at"`
	}
)
