package dto



type PemesananUpdateDTO struct {
	Id_pemesanan      uint64    `json:"id_pemesanan"`
	Id_customer       uint64    `json:"id_customer,omitempty" form:"id_customer,omitempty"`
	Tanggal_pemesanan string `json:"tanggal_pemesanan"`
	Total_pembayaran  uint64    `json:"total_pembayaran" form:"total_pembayaran" binding:"required"`
	Status            string    `json:"status" form:"status"`
	Bukti_Bayar            string    `json:"bukti_bayar" form:"bukti_bayar" binding:"required"`
}

type PemesananCreateDTO struct {
	Id_customer       uint64    `json:"id_customer,omitempty" form:"id_customer,omitempty"`
	Tanggal_pemesanan string `json:"tanggal_pemesanan"`
	Total_pembayaran  uint64    `json:"total_pembayaran" form:"total_pembayaran" binding:"required"`
	Status            string    `json:"status" form:"status"`
	Bukti_Bayar            string    `json:"bukti_bayar" form:"bukti_bayar" binding:"required"`
}
