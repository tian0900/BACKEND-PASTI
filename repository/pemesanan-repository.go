package repository

import (
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"gorm.io/gorm"
)

type PemesananRepository interface {
	InsertPemesanan(b entity.Pemesanan) entity.Pemesanan
	UpdatePemesanan(b entity.Pemesanan) entity.Pemesanan
	DeletePemesanan(b entity.Pemesanan)
	AllPemesanan() []entity.Pemesanan
	FindPemesananID(PemesananID uint64) entity.Pemesanan
}

type pemesananConnection struct {
	connection *gorm.DB
}

func NewPemesananRepository(dbConn *gorm.DB) PemesananRepository {
	return &pemesananConnection{
		connection: dbConn,
	}
}

func (db *pemesananConnection) InsertPemesanan(b entity.Pemesanan) entity.Pemesanan {
	db.connection.Save(&b)
	return b
}

func (db *pemesananConnection) UpdatePemesanan(b entity.Pemesanan) entity.Pemesanan {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *pemesananConnection) DeletePemesanan(b entity.Pemesanan){
	db.connection.Delete(&b)
}

func (db *pemesananConnection) FindPemesananID(pemesananID uint64) entity.Pemesanan {
	var pemesanan entity.Pemesanan
	db.connection.Find(&pemesanan, pemesananID)
	return pemesanan
}


func (db *pemesananConnection) AllPemesanan() []entity.Pemesanan {
	var pemesanans []entity.Pemesanan
	db.connection.Preload("User").Find(&pemesanans)
	return pemesanans
}
