package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
	"github.com/joho/godotenv"
	"github.com/worldofprasanna/fchat-server/models"
)

//ConnectDB function: Make database connection
func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("databaseUser")
	password := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	databaseHost := os.Getenv("databaseHost")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", databaseHost, username, databaseName, password)

	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		log.Fatal("error", err)
		panic(err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Message{},
	)

	log.Println("Successfully connected!", db)
	return db
}
