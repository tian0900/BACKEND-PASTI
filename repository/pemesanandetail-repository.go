package repository

import (
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"gorm.io/gorm"
)

type PemesananDetailRepository interface {
	InsertPemesananDetail(b entity.PemesananDetail) entity.PemesananDetail
	UpdatePemesananDetail(b entity.PemesananDetail) entity.PemesananDetail
	DeletePemesananDetail(b entity.PemesananDetail)
	AllPemesananDetail() []entity.PemesananDetail
	FindPemesananDetailID(PemesananDetailID uint64) entity.PemesananDetail
}

type pemesanandetailConnection struct {
	connection *gorm.DB
}

func NewPemesananDetailkRepository(dbConn *gorm.DB) PemesananDetailRepository {
	return &pemesanandetailConnection{
		connection: dbConn,
	}
}

func (db *pemesanandetailConnection) InsertPemesananDetail(b entity.PemesananDetail) entity.PemesananDetail {
	db.connection.Save(&b)
	return b
}

func (db *pemesanandetailConnection) UpdatePemesananDetail(b entity.PemesananDetail) entity.PemesananDetail {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *pemesanandetailConnection) DeletePemesananDetail(b entity.PemesananDetail){
	db.connection.Delete(&b)
}

func (db *pemesanandetailConnection) FindPemesananDetailID(pemesanandetailID uint64) entity.PemesananDetail {
	var pemesanandetail entity.PemesananDetail
	db.connection.Find(&pemesanandetail, pemesanandetailID)
	return pemesanandetail
}

func (db *pemesanandetailConnection) AllPemesananDetail() []entity.PemesananDetail {
	var pemesanandetails []entity.PemesananDetail
	db.connection.Preload("User").Find(&pemesanandetails)
	return pemesanandetails
}

