//The server package privdes the basic gofire server
package server

import (
	//"encoding/json"
	"gofire/user"
	"net/http"
	//"net/url"
	"net"
	//"time"
	//"strings"
)

//Rounting of the restful api
const (
	API      = "/api"
	CHAT     = "/api/c"
	CHATROOM = "/api/c/"
)

type FireServer struct {
	Addr                string      `json:"-"` //The Adress the server is running on
	RegisteredChatRooms []string    //All registered chatrooms
	User                []user.User `json:"-"` //All user on the chatroom
	CloseChannel        chan int
}

func addRestHandler(fireServer *FireServer, fn func(*FireServer, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		fn(fireServer, w, request)
	}
}

//a fireserver instance
var restCommands = make(map[string]string)

func init() {
	//AddRestCommand(API, ApiHandler, "Get all commands")
	//AddRestCommand(CHAT, ChatRoomHandler, "Get all chatrooms")
	//AddRestCommand(CHATROOM, SpecificChatRoomHandler, "Get specific chatroom info")
	//initServer()
	http.HandleFunc("/hello", HelloHandler)
}

//adds a rest command 
func AddRestCommand(pattern string, handler func(http.ResponseWriter, *http.Request), desc string) {
	restCommands[pattern] = desc
	http.HandleFunc(pattern, handler)
}

func (fireServer *FireServer) Close() {
	fireServer.CloseChannel <- 0
}

func (fireServer *FireServer) run() error{
	mux := http.NewServeMux()
	mux.HandleFunc("/api/addr", addRestHandler(fireServer, HelloAddress))
	mux.HandleFunc("/hello", HelloHandler)
	s := &http.Server{
		Handler:        mux,
	}
	l, e := net.Listen("tcp", fireServer.Addr)

	if e != nil {
		panic(e.Error())
	}

	go func(){
		<-fireServer.CloseChannel
		l.Close()
	}()

	return	s.Serve(l)
}

func HelloAddress(fireServer *FireServer, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fireServer.Addr))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

//A wrapper for the ListenAndServe of net/http
func NewFireServer(addr string) *FireServer{
	fireServer := &FireServer{
		Addr:                addr,
		RegisteredChatRooms: make([]string, 0),
		User:                make([]user.User, 0),
		CloseChannel:        make(chan int, 1),
	}
	return fireServer
}

func (fireServer *FireServer) ListenAndServe()error{
	return fireServer.run()
}

/*func ApiHandler(w http.ResponseWriter, r *http.Request) {
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
}*/
