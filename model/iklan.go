package model

type (
	Iklan struct {
		ID        int    `json:"id"`
		Thumbnail string `json:"thumbnail"`
		// CreatedAt time.Time `json:"created_at"`
		// UpdatedAt time.Time `json:"updated_at"`
	}
)
