//The server package privdes the basic gofire server
package server

import (
	"encoding/json"
	"gofire/user"
	"net/http"
	"net/url"
	"strings"
)

//Rounting of the restful api
const (
	API      = "/api"
	CHAT     = "/api/c"
	CHATROOM = "/api/c/"
)

type FireServer struct {
	Addr                string `json:"-"`
	RegisteredChatRooms []string
	User                []user.User `json:"-"`
}

//a fireserver instance
var fireServer = new(FireServer)
var restCommands = make(map[string]string)

func init() {
	AddRestCommand(API, ApiHandler, "Get all commands")
	AddRestCommand(CHAT, ChatRoomHandler, "Get all chatrooms")
	AddRestCommand(CHATROOM, SpecificChatRoomHandler, "Get specific chatroom info")
	//initServer()
}

//adds a rest command 
func AddRestCommand(pattern string, handler func(http.ResponseWriter, *http.Request), desc string) {
	restCommands[pattern] = desc
	http.HandleFunc(pattern, handler)
}

//A wrapper for the ListenAndServe of net/http
func ListenAndServe(addr string) error {
	fireServer.Addr = addr
	err := http.ListenAndServe(addr, nil)
	return err
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(restCommands)
	if err != nil {
		w.Write([]byte("404"))
	} else {
		w.Write(json)
	}
}

//ChatRoomHandler
func ChatRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getAllChatrooms(w, r)
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			//TODO Write error better
			w.Write([]byte([]byte(string(http.StatusBadRequest))))
		} else {
			name := r.FormValue("name")
			fireServer.RegisteredChatRooms = append(fireServer.RegisteredChatRooms, name)
			w.Write([]byte(name))
		}
	}
}

func postChatRoom(form url.Values, w http.ResponseWriter) {
	name := form.Get("name")
	fireServer.RegisteredChatRooms = append(fireServer.RegisteredChatRooms, name)
	w.Write([]byte(string(http.StatusOK)))
}

// get /c 
func getAllChatrooms(w http.ResponseWriter, r *http.Request) {
	if len(fireServer.RegisteredChatRooms) != 0 {
		json, err := json.Marshal(fireServer.RegisteredChatRooms)
		if err == nil {
			w.Write(json)
		} else {
			w.Write([]byte(string(http.StatusNotFound)))
		}

	} else {
		w.Write([]byte(string(http.StatusNotFound)))
	}
}

//get information about an specific chatroom
func SpecificChatRoomHandler(w http.ResponseWriter, r *http.Request) {
	name := getChatRoomName(r.URL.Path)
	if name != "" {
		roomcommand := strings.Split(name, "/")
		w.Write([]byte(roomcommand[1]))
	} else {
		w.Write([]byte(string(http.StatusBadRequest)))
	}
}

func isCommand(input string) bool {
	return false
}

//helper function for getting the chatroom
func getChatRoomName(link string) string {
	if len(link) == len(CHATROOM) {
		return ""
	}

	return link[len(CHATROOM):]
}
