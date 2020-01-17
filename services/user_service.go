package services

import (
	"github.com/jinzhu/gorm"
	"github.com/worldofprasanna/fchat-server/models"
)

// UserService : interface to create user and get all users
type UserService interface {
	CreateUser(user *models.User) *gorm.DB
	AllUsers() *gorm.DB
}

type userService struct {
	db *gorm.DB
}

// NewUserService - constructor to create new USPServerAPI
func NewUserService(db *gorm.DB) UserService {
	return userService{
		db: db,
	}
}

func (us userService) CreateUser(user *models.User) *gorm.DB {
	return us.db.Create(&user)
}

func (us userService) AllUsers() *gorm.DB {
	return models.AllUsers(us.db)
}
