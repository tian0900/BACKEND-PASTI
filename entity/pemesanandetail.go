package entity

type PemesananDetail struct {
	Id_pemesanandetail uint64 `gorm:"primary_key:auto_increment"`
	Id_pemesanan       uint64 `gorm:"type:integer" json:"id_pemesanan"`
	Id_produk          uint64 `gorm:"type:integer" json:"id_produk"`
	Kuantitas_pesan    uint64 `gorm:"type:integer" json:"kuantitas_pesan"`
	Total_harga        uint64 `gorm:"type:integer" json:"total_harga"`

}
