package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/worldofprasanna/fchat-server/routes"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(e)

	port := os.Getenv("PORT")

	http.Handle("/", routes.Handlers())

	log.Printf("Server up and running on port ':%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
