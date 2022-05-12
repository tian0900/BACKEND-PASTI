package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/dto"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/repository"
)

type KeranjangService interface {
	Insert(b dto.KeranjangCreateDTO) entity.Keranjang
	Update(b dto.KeranjangUpdateDTO) entity.Keranjang
	Delete(b entity.Keranjang)
	All() []entity.Keranjang
	FIndById(KeranjangID uint64) entity.Keranjang
	FIndByUserId(id_customer  uint64) []entity.Keranjang
	IsAllowedToEdit(userID string, KeranjangID uint64) bool
}

type keranjangService struct {
	keranjangRepository repository.KeranjangRepository
	userRepository repository.UserRepository
}



func NewKeranjangService(keranjangRepo repository.KeranjangRepository) KeranjangService {
	return &keranjangService{
		keranjangRepository: keranjangRepo,
	}
}

func (service *keranjangService) Insert(b dto.KeranjangCreateDTO) entity.Keranjang { //menambah data pada buku 
	keranjang := entity.Keranjang{}
	err := smapping.FillStruct(&keranjang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.keranjangRepository.InsertKeranjang(keranjang)
	return res
}

func (service *keranjangService) Update(b dto.KeranjangUpdateDTO) entity.Keranjang { //update data pada buku 
	keranjang := entity.Keranjang{}
	err := smapping.FillStruct(&keranjang, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.keranjangRepository.UpdateKeranjang(keranjang)
	return res
}

func (service *keranjangService) Delete(b entity.Keranjang) { //hapus data buku 
	service.keranjangRepository.DeleteKeranjang(b)
}

func (service *keranjangService) All() []entity.Keranjang {	//tampil semua buku 
	return service.keranjangRepository.AllKeranjang()
}

func (service *keranjangService) FIndById(keranjangID uint64) entity.Keranjang{ //menemukan buku sesuai ID
	return service.keranjangRepository.FindKeranjangID(keranjangID)
}

func (service *keranjangService) FIndByUserId(keranjangID uint64) []entity.Keranjang{ //menemukan buku sesuai ID
	return service.keranjangRepository.FindByUserID(keranjangID)
}


func (service *keranjangService) IsAllowedToEdit(userID string, keranjangID uint64) bool { 
	b := service.userRepository.FindByID(userID)
	id := fmt.Sprintf("%v", b.User_id)
	return userID == id
}
