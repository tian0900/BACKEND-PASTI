package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/dto"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/repository"
)

type ProdukService interface {
	Insert(b dto.ProdukCreateDTO) entity.Produk
	Update(b dto.ProdukUpdateDTO) entity.Produk
	Delete(b entity.Produk)
	All() []entity.Produk
	FIndById(produkID uint64) entity.Produk
	IsAllowedToEdit(userID string, produkID uint64) bool
}

type produkService struct {
	produkRepository repository.ProdukRepository
	userRepository repository.UserRepository
}



func NewProdukService(produkRepo repository.ProdukRepository) ProdukService {
	return &produkService{
		produkRepository: produkRepo,
	}
}

func (service *produkService) Insert(b dto.ProdukCreateDTO) entity.Produk { //menambah data pada buku 
	produk := entity.Produk{}
	err := smapping.FillStruct(&produk, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.produkRepository.InsertProduk(produk)
	return res
}

func (service *produkService) Update(b dto.ProdukUpdateDTO) entity.Produk { //update data pada buku 
	produk := entity.Produk{}
	err := smapping.FillStruct(&produk, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.produkRepository.UpdateProduk(produk)
	return res
}

func (service *produkService) Delete(b entity.Produk) { //hapus data buku 
	service.produkRepository.DeleteProduk(b)
}

func (service *produkService) All() []entity.Produk {	//tampil semua buku 
	return service.produkRepository.AllProduk()
}

func (service *produkService) FIndById(produkID uint64) entity.Produk { //menemukan buku sesuai ID
	return service.produkRepository.FindProdukID(produkID)
}

func (service *produkService) IsAllowedToEdit(userID string, produkID uint64) bool { 
	b := service.userRepository.FindByID(userID)
	id := fmt.Sprintf("%v", b.User_id)
	return userID == id
}
