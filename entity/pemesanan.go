package entity

type Pemesanan struct {
	Id_pemesanan      uint64    `gorm:"primary_key:auto_increment"`
	Id_customer       uint64    `gorm:"type:integer" json:"id_customer"`
	Tanggal_pemesanan string `gorm:"type:integer" json:"tanggal_pemesanan"`
	Total_pembayaran  uint64    `gorm:"type:integer" json:"total_pembayaran"`
	Status            string    `gorm:"type:string" json:"status"`
	Bukti_Bayar       string    `gorm:"type:varchar(255)" json:"bukti_bayar"`

	User User `gorm:"foreignkey:id_customer;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"`
}
