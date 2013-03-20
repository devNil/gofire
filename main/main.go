package main

import (
	//"code.google.com/p/go.net/websocket"
	"gofire/command"
	"gofire/message"
	"net/http"
	//"os"
	//"os/signal"
	srv "gofire/server"
	"gofire/server/api"
)

//constants 
const ADDR = ":8080"
const TEMPDIR = "temp/"
const STATICDIR = "template/"

//global vars 
//default chatroom
var chatRoom = ChatRoom{
	name:                  "unity is gay", //this is not meant to be offensive
	history:               make([]*message.Message, 0),
	broadcast:             make(chan *command.Command),
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

	//go server.run()

	//cleanup if command+c
	/*c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		server.cleanUp()
		os.Exit(0)
	}()*/



	fireserver := srv.NewFireServer()

	http.HandleFunc("/api",srv.AddHandleFunc(fireserver,api.Ping))

	//fireserver.ListenAndServe()

	//http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/index.html", loginHandler)
	//http.HandleFunc("/chat.html", chatHandler)
	//http.HandleFunc("/doLogin", doLogin)
	//http.Handle("/ws", websocket.Handler(wsHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
