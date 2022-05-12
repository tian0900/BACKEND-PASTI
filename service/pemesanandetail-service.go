package service

import (
	"fmt"
	"log"

	"github.com/NestyTampubolon/golang_gin_gorm_JWT/dto"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/repository"
	"github.com/mashingan/smapping"
)

type PemesananDetailService interface {
	Insert(b dto.PemesananDetailCreateDTO) entity.PemesananDetail
	All() []entity.PemesananDetail
	FIndById(pemesanandetailID uint64) entity.PemesananDetail
	IsAllowedToEdit(userID string, pemesanandetailID uint64) bool
}

type pemesanandetailService struct {
	pemesanandetailRepository repository.PemesananDetailRepository
	userRepository      repository.UserRepository
}

func NewPemesananDetailService(pemesanandetailRepo repository.PemesananDetailRepository) PemesananDetailService {
	return &pemesanandetailService{
		pemesanandetailRepository: pemesanandetailRepo,
	}
}

func (service *pemesanandetailService) Insert(b dto.PemesananDetailCreateDTO) entity.PemesananDetail { //menambah data pada buku
	pemesanandetail := entity.PemesananDetail{}
	err := smapping.FillStruct(&pemesanandetail, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.pemesanandetailRepository.InsertPemesananDetail(pemesanandetail)
	return res
}


func (service *pemesanandetailService) All() []entity.PemesananDetail { //tampil semua buku
	return service.pemesanandetailRepository.AllPemesananDetail()
}

func (service *pemesanandetailService) FIndById(pemesanandetailID uint64) entity.PemesananDetail { //menemukan buku sesuai ID
	return service.pemesanandetailRepository.FindPemesananDetailID(pemesanandetailID)
}

func (service *pemesanandetailService) IsAllowedToEdit(userID string, pemesanandetailID uint64) bool {
	b := service.userRepository.FindByID(userID)
	id := fmt.Sprintf("%v", b.User_id)
	return userID == id
}
