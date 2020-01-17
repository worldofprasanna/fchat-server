package services

import (
	"github.com/jinzhu/gorm"
	"github.com/worldofprasanna/fchat-server/models"
)

// MessageService :
type MessageService interface {
	CreateMessage(message *models.Message) *gorm.DB
	AllMessages(senderID, receiverID string) *gorm.DB
}

type messageService struct {
	db *gorm.DB
}

// NewMessageService - constructor to create new message service
func NewMessageService(db *gorm.DB) MessageService {
	return messageService{
		db: db,
	}
}

func (us messageService) CreateMessage(message *models.Message) *gorm.DB {
	return us.db.Create(&message)
}

func (us messageService) AllMessages(senderID, receiverID string) *gorm.DB {
	return models.AllMessages(us.db, senderID, receiverID)
}
