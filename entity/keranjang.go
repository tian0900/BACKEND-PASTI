package entity

type Keranjang struct {
	Id_keranjang uint64 `gorm:"primary_key:auto_increment"`
	Id_customer  uint64 `gorm:"type:integer" json:"id_customer"`
	Id_produk    uint64 `gorm:"type:integer" json:"id_produk"`
	Kuantitas    uint64 `gorm:"type:integer" json:"kuantitas"`
	Total        uint64 `gorm:"type:integer" json:"total"`

	User   User   `gorm:"foreignkey:id_customer;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"`
	Produk Produk `gorm:"foreignkey:Id_produk;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"produk"`
}
