package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/worldofprasanna/fchat-server/models"
	"github.com/worldofprasanna/fchat-server/services"
	"github.com/worldofprasanna/fchat-server/utils"
)

// Handlers function
func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	var db = utils.ConnectDB()
	userService := services.NewUserService(db)
	messageService := services.NewMessageService(db)

	fmt.Print(userService)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	})

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		user := models.NewUser(db, r.Body)
		createdUser := userService.CreateUser(user)
		json.NewEncoder(w).Encode(createdUser)
	}).Methods("POST")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := userService.AllUsers()
		json.NewEncoder(w).Encode(users)
	})

	r.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		message := models.NewMessage(db, r.Body)
		createdMessage := messageService.CreateMessage(message)
		json.NewEncoder(w).Encode(createdMessage)
	}).Methods("POST")

	r.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		allMessages := messageService.AllMessages(r.FormValue("sender_id"), r.FormValue("receiver_id"))
		json.NewEncoder(w).Encode(allMessages)
	})

	return r
}

// CommonMiddleware function to set content type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
