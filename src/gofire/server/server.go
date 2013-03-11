//The server package privdes the basic gofire server
package server

import (
	//"encoding/json"
	"gofire/user"
	"net/http"
	//"net/url"
	//"time"
	//"strings"
)

type FireServer struct {
	RegisteredChatRooms []string    //All registered chatrooms
	User                []user.User `json:"-"` //All user on the chatroom
}

type FireServerHandler func(*FireServer, http.ResponseWriter, *http.Request)

func AddHandler(fireServer *FireServer, fn FireServerHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		fn(fireServer, w, request)
	}
}

func NewFireServer() *FireServer{
	fireServer := &FireServer{
		RegisteredChatRooms: make([]string, 0),
		User:                make([]user.User, 0),
	}
	return fireServer
}
