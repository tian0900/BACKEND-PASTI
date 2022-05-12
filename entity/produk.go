package entity

type Produk struct {
	Produk_id uint64 `gorm:"primary_key:auto_increment" json:"produk_id"`
	Nama      string `gorm:"type:varchar(255)" json:"nama"`
	Harga     uint64 `gorm:"type:integer" json:"harga,string"`
	Gambar    string `gorm:"type:varchar(255)" json:"gambar"`
	Stok      uint64 `gorm:"type:integer" json:"stok,string"`
	Kategori  string `gorm:"type:varchar(255)" json:"kategori"`
}
