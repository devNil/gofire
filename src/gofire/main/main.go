package main

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
	"os"
	"os/signal"
)

//constants 
const ADDR = ":8080"
const TEMPDIR = "temp/"
const STATICDIR = "template/"

type Message struct {
	Usr *User
	Msg []byte
}

//global vars 
//default chatroom
var chatRoom = ChatRoom{
	name:                  "unity is gay",
	history:               make([]*Message, 0),
	broadcast:             make(chan *Message),
	register:              make(chan *Connection),
	unregister:            make(chan *Connection),
	registeredConnections: make(map[*Connection]bool),
}
var server = Server{
	chatRooms:             make([]*ChatRoom, 0),
	register:              make(chan *Connection),
	unregister:            make(chan *Connection),
	registeredConnections: make(map[*Connection]bool),
}

func main() {

	go server.run()

	//cleanup if command+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		server.cleanUp()
		os.Exit(0)
	}()

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/index.html", loginHandler)
	http.HandleFunc("/chat.html", chatHandler)
	http.HandleFunc("/doLogin", doLogin)
	http.Handle("/ws", websocket.Handler(wsHandler))
	err := http.ListenAndServe(ADDR, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
