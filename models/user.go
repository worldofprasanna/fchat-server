package models

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model

	UserName string `gorm:"type:varchar(100);unique_index"`
}
