//This package provides all rest apis
package api

import(
	"gofire/server"
	"net/http"
)

const(
	POST = "POST"
	GET = "GET"
)

//differs POST or GET /api/c
func ChatRoomHandler(fireServer *server.FireServer, w http.ResponseWriter, r *http.Request){
	if r.Method == POST{
		postChatRoomHandler(fireServer, w, r)
	}
	if r.Method == GET{
		getChatRoomHandler(fireServer, w, r)
	}
}

func postChatRoomHandler(fireServer *server.FireServer, w http.ResponseWriter, r *http.Request){

}

func getChatRoomHandler(fireServer *server.FireServer, w http.ResponseWriter, r *http.Request){

}
