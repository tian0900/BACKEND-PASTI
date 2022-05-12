package repository

import (
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"gorm.io/gorm"
)

type KeranjangRepository interface {
	InsertKeranjang(b entity.Keranjang) entity.Keranjang
	UpdateKeranjang(b entity.Keranjang) entity.Keranjang
	DeleteKeranjang(b entity.Keranjang)
	AllKeranjang() []entity.Keranjang
	FindKeranjangID(KeranjangID uint64) entity.Keranjang
	FindByUserID(KeranjangID uint64) []entity.Keranjang
}

type keranjangConnection struct {
	connection *gorm.DB
}

func NewKeranjangRepository(dbConn *gorm.DB) KeranjangRepository {
	return &keranjangConnection{
		connection: dbConn,
	}
}

func (db *keranjangConnection) InsertKeranjang(b entity.Keranjang) entity.Keranjang {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *keranjangConnection) UpdateKeranjang(b entity.Keranjang) entity.Keranjang {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *keranjangConnection) DeleteKeranjang(b entity.Keranjang) {
	db.connection.Delete(&b)
}

func (db *keranjangConnection) FindKeranjangID(KeranjangID uint64) entity.Keranjang {
	var keranjang entity.Keranjang
	db.connection.Find(&keranjang, KeranjangID)
	return keranjang
}

func (db *keranjangConnection) FindByUserID(id_customer uint64) []entity.Keranjang {
	var user []entity.Keranjang
	db.connection.Where("id_customer = ?", id_customer).Preload("User").Find(&user)
	return user
}

func (db *keranjangConnection) AllKeranjang() []entity.Keranjang {
	var keranjangs []entity.Keranjang
	db.connection.Preload("User").Find(&keranjangs)
	return keranjangs
}
