package repository

import (
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"gorm.io/gorm"
)

type ProdukRepository interface {
	InsertProduk(b entity.Produk) entity.Produk
	UpdateProduk(b entity.Produk) entity.Produk
	DeleteProduk(b entity.Produk)
	AllProduk() []entity.Produk
	FindProdukID(ProdukID uint64) entity.Produk
}

type produkConnection struct {
	connection *gorm.DB
}

func NewProdukRepository(dbConn *gorm.DB) ProdukRepository {
	return &produkConnection{
		connection: dbConn,
	}
}

func (db *produkConnection) InsertProduk(b entity.Produk) entity.Produk {
	db.connection.Save(&b)
	return b
}

func (db *produkConnection) UpdateProduk(b entity.Produk) entity.Produk {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *produkConnection) DeleteProduk(b entity.Produk){
	db.connection.Delete(&b)
}

func (db *produkConnection) FindProdukID(produkID uint64) entity.Produk {
	var produk entity.Produk
	db.connection.Find(&produk, produkID)
	return produk
}

func (db *produkConnection) AllProduk() []entity.Produk {
	var produks []entity.Produk
	db.connection.Preload("User").Find(&produks)
	return produks
}
