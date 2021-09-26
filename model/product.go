package model

type (
	Product struct {
		ID           int    `json:"id"`
		Toko_Id      int    `json:"toko_id"`
		Kategori_Id  int    `json:"kategori_id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Harga        int    `json:"harga"`
		ImageUrl     string `json:"image_url"`
		StatusDelete string `json:"statusdelete"`
		// CreatedAt    time.Time `json:"created_at"`
		// UpdatedAt    time.Time `json:"updated_at"`
	}
)
