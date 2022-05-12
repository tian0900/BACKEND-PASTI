package dto

type PemesananDetailCreateDTO struct {
	Id_pemesanan    uint64 `json:"id_pemesanan,omitempty" form:"id_pemesanan,omitempty"`
	Id_produk       uint64 `json:"id_produk,omitempty" form:"id_produk,omitempty"`
	Kuantitas_pesan uint64 `json:"kuantitas_pesan"`
	Total_harga     uint64 `json:"total_Total_harga" form:"total_Total_harga" binding:"required"`
}
