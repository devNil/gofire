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

type FireServerFunc func(*FireServer, http.ResponseWriter, *http.Request)

type FireServerHandler interface{
	Handle(*FireServer, http.ResponseWriter, *http.Request)
}

func AddHandleFunc(fireServer *FireServer, fn FireServerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		fn(fireServer, w, request)
	}
}

func AddHandler(fireServer *FireServer, handler FireServerHandler) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, request *http.Request){
		handler.Handle(fireServer, w, request)
	}

}

func NewFireServer() *FireServer{
	fireServer := &FireServer{
		RegisteredChatRooms: make([]string, 0),
		User:                make([]user.User, 0),
	}
	return fireServer
}
