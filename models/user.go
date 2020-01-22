package models

import (
	"encoding/json"
	"io"

	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model

	UserName string `gorm:"type:varchar(100);unique_index"`
}

// NewUser - create a new user in db
func NewUser(db *gorm.DB, data io.ReadCloser) *User {
	var user User
	json.NewDecoder(data).Decode(&user)
	return &user
}

// AllUsers - get all users from db
func AllUsers(db *gorm.DB) *gorm.DB {
	var users []User
	return db.Find(&users)
}
