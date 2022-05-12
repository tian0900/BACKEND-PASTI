package service

import (
	"fmt"
	"log"

	"github.com/NestyTampubolon/golang_gin_gorm_JWT/dto"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/repository"
	"github.com/mashingan/smapping"
)

type FeedbackService interface {
	Insert(b dto.FeedbackCreateDTO) entity.Feedback
	Update(b dto.FeedbackUpdateDTO) entity.Feedback
	Delete(b entity.Feedback)
	All() []entity.Feedback
	FIndById(FeedbackID uint64) entity.Feedback
	IsAllowedToEdit(userID string, FeedbackID uint64) bool
}

type feedbackService struct {
	feedbackRepository repository.FeedbackRepository
	userRepository     repository.UserRepository
}

func NewFeedbackService(feedbackRepo repository.FeedbackRepository) FeedbackService {
	return &feedbackService{
		feedbackRepository: feedbackRepo,
	}
}

func (service *feedbackService) Insert(b dto.FeedbackCreateDTO) entity.Feedback { //menambah data pada buku
	feedback := entity.Feedback{}
	err := smapping.FillStruct(&feedback, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.feedbackRepository.InsertFeedback(feedback)
	return res
}

func (service *feedbackService) Update(b dto.FeedbackUpdateDTO) entity.Feedback { //update data pada buku
	feedback := entity.Feedback{}
	err := smapping.FillStruct(&feedback, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.feedbackRepository.UpdateFeedback(feedback)
	return res
}

func (service *feedbackService) Delete(b entity.Feedback) { //hapus data buku
	service.feedbackRepository.DeleteFeedback(b)
}

func (service *feedbackService) All() []entity.Feedback { //tampil semua buku
	return service.feedbackRepository.AllFeedback()
}

func (service *feedbackService) FIndById(feedbackID uint64) entity.Feedback { //menemukan buku sesuai ID
	return service.feedbackRepository.FindFeedbackID(feedbackID)
}

func (service *feedbackService) IsAllowedToEdit(userID string, feedbackID uint64) bool {
	b := service.userRepository.FindByID(userID)
	id := fmt.Sprintf("%v", b.User_id)
	return userID == id
}
