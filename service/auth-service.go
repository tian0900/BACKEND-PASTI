package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/dto"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/repository"
	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface { //fungsi yang ada pada Auth
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} { ///verifikasi email dan password
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparePassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparePassword {
			return res
		}
		return false
	}
	return res
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User { //menambah user
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) entity.User { //menemukan data berdasarkan email
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool { //check email yang duplicate
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false
	}
	return true
}
