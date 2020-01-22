package models

import (
	"encoding/json"
	"io"

	"github.com/jinzhu/gorm"
)

// Message struct
type Message struct {
	gorm.Model

	Text     string `json:"text"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

// NewMessage - create message in db
func NewMessage(db *gorm.DB, data io.ReadCloser) *Message {
	var message Message
	json.NewDecoder(data).Decode(&message)
	return &message
}

// AllMessages - get all users from db
func AllMessages(db *gorm.DB, sender, receiver string) *gorm.DB {
	var messages []Message
	query := "(sender = ? AND receiver = ? ) OR (receiver = ? AND sender = ? )"
	return db.Where(query, sender, receiver, sender, receiver).Find(&messages)
}
