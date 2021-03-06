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

const StandardAddress = "localhost"

func main(){

	addr := os.Getenv("ADRESS")

	if addr == ""{
		addr = StandardAddress
	}

	port := os.Getenv("PORT")

	if port == ""{
		port = StandardPort
	}

	http.HandleFunc("/", web.IndexHandler)
	log.Println("IndexHandler registered")

    http.HandleFunc("/css/", web.StaticHandler)
    http.HandleFunc("/img/", web.StaticHandler)

	http.HandleFunc("/login", web.LoginHandler)
	log.Println("LoginHandler registered")

	http.HandleFunc("/logout", web.LogoutHandler)
	log.Println("LogoutHandler registered")

	http.HandleFunc("/chat", web.ChatHandler)
	log.Println("ChatHandler registered")

	socket.Start()
	log.Println("Fireserver is running")

	http.Handle("/ws",websocket.Handler(socket.SocketHandler))

	log.Printf("Server started on : %s:%s",addr, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s",addr,port),nil)
	panic(err)
}
