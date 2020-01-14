package controllers

import (
	"encoding/json"
	"github.com/worldofprasanna/fchat-server/models"
	"github.com/worldofprasanna/fchat-server/utils"
	"net/http"
)

var db = utils.ConnectDB()

// TestAPI function
func TestAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is rocking!"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	createdUser := db.Create(user)
	json.NewEncoder(w).Encode(createdUser)
}
