package dto

import "time"

type PemesananUpdateDTO struct {
	Id_pemesanan      uint64    `json:"id_pemesanan"`
	Id_customer       uint64    `json:"id_customer,omitempty" form:"id_customer,omitempty"`
	Tanggal_pemesanan time.Time `json:"tanggal_pemesanan"`
	Total_pembayaran  uint64    `json:"total_pembayaran" form:"total_pembayaran" binding:"required"`
	Status            string    `json:"status" form:"status"`
	Alamat            string    `json:"alamat" form:"alamat" binding:"required"`
}

type PemesananCreateDTO struct {
	Id_customer       uint64    `json:"id_customer,omitempty" form:"id_customer,omitempty"`
	Tanggal_pemesanan time.Time `json:"tanggal_pemesanan"`
	Total_pembayaran  uint64    `json:"total_pembayaran" form:"total_pembayaran" binding:"required"`
	Status            string    `json:"status" form:"status"`
	Alamat            string    `json:"alamat" form:"alamat" binding:"required"`
}
