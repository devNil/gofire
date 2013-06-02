package web

import(
	db "gofire/database"
	"code.google.com/p/go.net/websocket"
	"log"
    "net/http"
    "gofire/socket"
)

func SocketHandler(conn *websocket.Conn){

    session, err := store.Get(conn.Request(), cookieName)

    if session.IsNew || err != nil{
        log.Println("Error Socket: ", err)
        http.Redirect(nil, conn.Request(), "/", http.StatusFound)
        return
    }

    id, _ := session.Values["id"].(int64)

    user, err := db.GetUser(id)

	if err != nil{
        log.Println("DB-Error: ",err)
		return
	}

    socket.RegisterConnection(conn, user)
}
