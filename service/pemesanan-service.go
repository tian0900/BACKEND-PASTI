package service

import (
	"fmt"
	"log"

	"github.com/NestyTampubolon/golang_gin_gorm_JWT/dto"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/repository"
	"github.com/mashingan/smapping"
)

type PemesananService interface {
	Insert(b dto.PemesananCreateDTO) entity.Pemesanan
	Update(b dto.PemesananUpdateDTO) entity.Pemesanan
	Delete(b entity.Pemesanan)
	All() []entity.Pemesanan
	FIndById(pemesananID uint64) entity.Pemesanan
	IsAllowedToEdit(userID string, pemesananID uint64) bool
}

type pemesananService struct {
	pemesananRepository repository.PemesananRepository
	userRepository      repository.UserRepository
}

func NewPemesananService(pemesananRepo repository.PemesananRepository) PemesananService {
	return &pemesananService{
		pemesananRepository: pemesananRepo,
	}
}

func (service *pemesananService) Insert(b dto.PemesananCreateDTO) entity.Pemesanan { //menambah data pada buku
	pemesanan := entity.Pemesanan{}
	err := smapping.FillStruct(&pemesanan, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.pemesananRepository.InsertPemesanan(pemesanan)
	return res
}

func (service *pemesananService) Update(b dto.PemesananUpdateDTO) entity.Pemesanan { //update data pada buku
	pemesanan := entity.Pemesanan{}
	err := smapping.FillStruct(&pemesanan, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.pemesananRepository.UpdatePemesanan(pemesanan)
	return res
}

func (service *pemesananService) Delete(b entity.Pemesanan) { //hapus data buku
	service.pemesananRepository.DeletePemesanan(b)
}

func (service *pemesananService) All() []entity.Pemesanan { //tampil semua buku
	return service.pemesananRepository.AllPemesanan()
}

func (service *pemesananService) FIndById(pemesananID uint64) entity.Pemesanan { //menemukan buku sesuai ID
	return service.pemesananRepository.FindPemesananID(pemesananID)
}

func (service *pemesananService) IsAllowedToEdit(userID string, pemesananID uint64) bool {
	b := service.userRepository.FindByID(userID)
	id := fmt.Sprintf("%v", b.User_id)
	return userID == id
}
