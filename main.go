package main

import(
	"fmt"
	"log"
	"net/http"
	"os"
	"gofire/web"
	"gofire/socket"
	"code.google.com/p/go.net/websocket"
)

const StandardPort = "8080"

func main(){

	port := os.Getenv("PORT")

	if port == ""{
		port = StandardPort
	}

	http.HandleFunc("/", web.IndexHandler)
	log.Println("IndexHandler registered")

	http.HandleFunc("/login", web.LoginHandler)
	log.Println("LoginHandler registered")

	http.HandleFunc("/chat", web.ChatHandler)
	log.Println("ChatHandler registered")

	socket.Start()
	log.Println("Fireserver is running")

	http.Handle("/ws",websocket.Handler(socket.SocketHandler))

	log.Printf("Server started on port: :%s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port),nil)
	panic(err)
}
