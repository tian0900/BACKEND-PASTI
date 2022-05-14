package dto

type ProdukUpdateDTO struct {
	Produk_id uint64 `json:"produk_id"`
	Nama      string `json:"nama" form:"nama"`
	Harga     uint64 `json:"harga,string" form:"harga"`
	Gambar    string `json:"gambar" form:"gambar"`
	Stok      uint64 `json:"stok,string" form:"stok"`
	Kuantitas uint64 `json:"kuantitas,string" form:"kuantitas"`
	Kategori  string `json:"kategori" form:"kategori"`
}

type ProdukCreateDTO struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Harga    uint64 `json:"harga,string" form:"harga" binding:"required"`
	Gambar   string `json:"gambar" form:"gambar" binding:"required"`
	Stok     uint64 `json:"stok,string" form:"stok" binding:"required"`
	Kategori string `json:"kategori" form:"kategori" binding:"required"`
}
