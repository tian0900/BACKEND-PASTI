package repository

import (
	"github.com/NestyTampubolon/golang_gin_gorm_JWT/entity"
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	InsertFeedback(b entity.Feedback) entity.Feedback
	UpdateFeedback(b entity.Feedback) entity.Feedback
	DeleteFeedback(b entity.Feedback)
	AllFeedback() []entity.Feedback
	FindFeedbackID(FeedbackID uint64) entity.Feedback
}

type feedbackConnection struct {
	connection *gorm.DB
}

func NewFeedbackRepository(dbConn *gorm.DB) FeedbackRepository {
	return &feedbackConnection{
		connection: dbConn,
	}
}

func (db *feedbackConnection) InsertFeedback(b entity.Feedback) entity.Feedback {
	db.connection.Save(&b)
	return b
}

func (db *feedbackConnection) UpdateFeedback(b entity.Feedback) entity.Feedback {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *feedbackConnection) DeleteFeedback(b entity.Feedback) {
	db.connection.Delete(&b)
}

func (db *feedbackConnection) FindFeedbackID(feedbackID uint64) entity.Feedback {
	var feedback entity.Feedback
	db.connection.Find(&feedback, feedbackID)
	return feedback
}

func (db *feedbackConnection) AllFeedback() []entity.Feedback {
	var feedbacks []entity.Feedback
	db.connection.Preload("User").Find(&feedbacks)
	return feedbacks
}
