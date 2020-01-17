package models

import (
	"encoding/json"
	"io"

	"github.com/jinzhu/gorm"
)

// Message struct
type Message struct {
	gorm.Model

	Text       string `json:"text"`
	SenderID   uint   `json:"sender_id"`
	ReceiverID uint   `json:"receiver_id"`

	Sender   User `gorm:"foreignkey:SenderID"`   // use SenderID as foreign key
	Receiver User `gorm:"foreignkey:ReceiverID"` // use ReceiverID as foreign key
}

// NewMessage - create message in db
func NewMessage(db *gorm.DB, data io.ReadCloser) *Message {
	var message Message
	json.NewDecoder(data).Decode(&message)
	return &message
}

// AllMessages - get all users from db
func AllMessages(db *gorm.DB, sID, rID string) *gorm.DB {
	var messages []Message
	query := "(sender_id = ? AND receiver_id = ? ) OR (receiver_id = ? AND sender_id = ? )"
	return db.Where(query, sID, rID, sID, rID).Find(&messages)
}
