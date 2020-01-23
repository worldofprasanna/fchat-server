package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/worldofprasanna/fchat-server/models"
	"github.com/worldofprasanna/fchat-server/services"
	"github.com/worldofprasanna/fchat-server/utils"
)

var clients = make(map[string]*websocket.Conn) // connected clients
var broadcast = make(chan models.Message)      // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// UserType : struct to distinguish socket response
type UserType struct {
	User models.User
	Type string
}

// MessageType : struct to distinguish socket response
type MessageType struct {
	Message models.Message
	Type    string
}

// UserList struct declaration
type UserList struct {
	Users      *gorm.DB
	ActiveList map[string]*websocket.Conn
	Type       string
}

// LogoutUser struct declaration
type LogoutUser struct {
	UserName string
}

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
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		user := models.User{}
		err = ws.ReadJSON(&user)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}
		_, err = userService.CreateUser(&user)
		if err != nil {
			log.Fatal(err)
		}
		for client := range clients {
			fmt.Println("Sending res to:", client)
			err := clients[client].WriteJSON(&UserList{Type: "UserList"})
			if err != nil {
				log.Printf("error: %v", err)
			}
		}

		clients[user.UserName] = ws
		fmt.Printf("all clients %+v", clients)

		userType := UserType{User: user, Type: "User"}

		if err = ws.WriteJSON(userType); err != nil {
			fmt.Println(err)
		}
	})

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := userService.AllUsers()
		json.NewEncoder(w).Encode(&UserList{Users: users, ActiveList: clients})
	})

	r.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		allMessages := messageService.AllMessages(r.FormValue("sender"), r.FormValue("receiver"))
		json.NewEncoder(w).Encode(allMessages)
	})

	// socket handler
	http.HandleFunc("/send_message", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		message := models.Message{}
		err = ws.ReadJSON(&message)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		messageService.CreateMessage(&message)
		messageType := MessageType{Message: message, Type: "Message"}
		receiverClient := clients[message.Receiver]
		if err = receiverClient.WriteJSON(messageType); err != nil {
			fmt.Println(err)
		}
		senderClient := clients[message.Sender]
		if err = senderClient.WriteJSON(messageType); err != nil {
			fmt.Println(err)
		}
	})

	r.HandleFunc("/logout/{UserName}", func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["UserName"]
		defer clients[name].Close()
		delete(clients, name)

		for client := range clients {
			fmt.Println("Sending res to:", client)
			err := clients[client].WriteJSON(&UserList{Type: "UserList"})
			if err != nil {
				log.Printf("error: %v", err)
			}
		}
		w.Write([]byte("logout_success"))
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
