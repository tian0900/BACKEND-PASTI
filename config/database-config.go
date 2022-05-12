package config

import (
	"fmt"
	"os"

	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB { //koneksi dengan database
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	dbUser := os.Getenv("DB_USER") //nama user database
	dbPass := os.Getenv("DB_PASS") //password
	dbHost := os.Getenv("DB_HOST") //nama host
	dbName := os.Getenv("DB_NAME") //nama database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(&entity.Pemesanan{}, &entity.PemesananDetail{}, &entity.Produk{}, &entity.User{}, &entity.Keranjang{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) { //menutup koneksi dengan database
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close a connection to database")
	}
	dbSQL.Close()
}
