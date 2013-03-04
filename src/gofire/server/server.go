//The server package privdes the basic gofire server
package server

import(
	"net/http"
	"encoding/json"
)

type FireServer struct{
	Name string
	Addr string
	RegisteredChatRooms []string
}

func (fs *FireServer)Init(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){fs.MainHandler(w, r)})
	http.HandleFunc(CHAT, func(w http.ResponseWriter, r *http.Request){fs.ChatRoomHandler(w, r)})
	err := http.ListenAndServe(fs.Addr, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (fs *FireServer)MainHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path[len(r.URL.Path)-3:] == "api" {
		fs.apiHandler(w, r)
	}
}

const CHAT = "/c/"
//ChatRoomHandler
func (fs *FireServer)ChatRoomHandler(w http.ResponseWriter, r *http.Request){
	chatName := r.URL.Path[len(CHAT):]
	if len(chatName) == 0{
		fs.getAllChatrooms(w, r)
	}else{
		w.Write([]byte(chatName))
	}
}

func (fs *FireServer)getAllChatrooms(w http.ResponseWriter, r *http.Request){
	json,err := json.Marshal(fs.RegisteredChatRooms)
	if err == nil{
		w.Write(json)
	}else{
		w.Write([]byte("Fail"))
	}
}

//Api Hanlder, Handles /api calls on server
func (fs *FireServer)apiHandler(w http.ResponseWriter, r *http.Request){
	json, err := json.Marshal(fs)
	if err != nil{
		w.Write([]byte("Fail"))
	}else{
		w.Write(json)
	}
}
