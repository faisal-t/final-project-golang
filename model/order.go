package model

type (
	Order struct {
		ID               int    `json:"id"`
		User_Id          int    `json:"user_id"`
		Produk_Id        string `json:"product_id"`
		Status           string `json:"status"`
		AlamatPengiriman string `json:"alamat_pengiriman"`
		Harga            int    `json:"harga"`
		Qty              int    `json:"qty"`
		TotalBayar       int    `json:"total_bayar"`
		BukTiBayar       string `json:"bukti_bayar"`
		// CreatedAt        time.Time `json:"created_at"`
		// UpdatedAt        time.Time `json:"updated_at"`
	}
)
