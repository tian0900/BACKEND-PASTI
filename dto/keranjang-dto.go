package dto

type KeranjangUpdateDTO struct {
	Id_keranjang uint64 `json:"id_keranjang"`
	Id_customer  uint64 `json:"id_customer,omitempty" form:"id_customer,omitempty"`
	Id_produk    uint64 `json:"id_produk"`
	Kuantitas    uint64 `json:"kuantitas" form:"kuantitas" binding:"required"`
	Total        uint64 `json:"total" form:"total" binding:"required"`
}

type KeranjangCreateDTO struct {
	Id_customer uint64 `json:"id_customer,omitempty" form:"id_customer,omitempty"`
	Id_produk   uint64 `json:"id_produk,string"`
	Harga       uint64 `json:"harga,string" form:"harga" binding:"required"`
	Stok        uint64 `json:"stok,string" form:"stok" binding:"required"`
	Kuantitas   uint64 `json:"kuantitas" form:"kuantitas"`
	Total       uint64 `json:"total" form:"total"`
}
