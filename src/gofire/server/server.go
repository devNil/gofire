//The server package privdes the basic gofire server
package server

import (
	"encoding/json"
	"net/http"
)

type FireServer struct {
	Name                string
	Addr                string `json:"-"`
	RegisteredChatRooms []string
}

//a fireserver instance
var fireServer = new(FireServer)
var restCommands = make(map[string]string)

func init() {
	AddRestCommand("/api", MainHandler, "Get all commands")
	AddRestCommand("/api/c", ChatRoomHandler, "Get all chatrooms")

	initServer()
}

//adds a rest command 
func AddRestCommand(pattern string, handler func(http.ResponseWriter, *http.Request), desc string) {
	restCommands[pattern] = desc
	http.HandleFunc(pattern, handler)
}

func ListenAndServe(addr string) error {
	fireServer.Addr = addr
	err := http.ListenAndServe(addr, nil)
	return err
}

func initServer() {
	http.HandleFunc("/", MainHandler)
	http.HandleFunc(CHAT, ChatRoomHandler)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[len(r.URL.Path)-3:] == "api" {
		apiHandler(w, r)
	}
}

//Chat-rounting
const CHAT = "/c/"

//ChatRoomHandler
func ChatRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		chatName := r.URL.Path[len(CHAT):]
		if len(chatName) == 0 {
			getAllChatrooms(w, r)
		} else {
			w.Write([]byte(chatName))
		}
	}

	if r.Method == "POST" {
		w.Write([]byte("Post to /c/"))
	}
}

// get /c/ 
func getAllChatrooms(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(fireServer.RegisteredChatRooms)
	if err == nil {
		w.Write(json)
	} else {
		w.Write([]byte("Fail"))
	}
}

func addChatRoom() {

}

//Api Hanlder, Handles /api calls on server
func apiHandler(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(restCommands)
	if err != nil {
		w.Write([]byte("Fail"))
	} else {
		w.Write(json)
	}
}
