package model

type (
	Toko struct {
		ID      int    `json:"id"`
		User_ID int    `json:"user_id"`
		Name    string `json:"name"`
		Alamat  string `json:"alamat"`
		// CreatedAt time.Time `json:"created_at"`
		// UpdatedAt time.Time `json:"updated_at"`
	}
)
