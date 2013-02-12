package main

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
)

//constants 
const ADDR = ":8080"
const TEMPDIR = "temp/"
const STATICDIR = "template/"

const (
	REGISTER = iota //0
	MESSAGE         //1
	BLOGIN
	BLOGOUT
)

//structs 
type User struct {
	Name string
}

type Command struct {
	Type  int
	Value []byte
}

type Connection struct {
	//User who sends
	Usr *User
	//websocket-connection
	Conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan *Message
}

type Message struct {
	Usr *User
	Msg []byte
}

//global vars 
var server = Server{
	history:               make([]*Message, 0),
	broadcast:             make(chan *Message),
	register:              make(chan *Connection),
	unregister:            make(chan *Connection),
	registeredConnections: make(map[*Connection]bool),
}

func (c *Connection) Read() {
	for {
		var message string
		err := websocket.Message.Receive(c.Conn, &message)
		if err != nil {
			break
		}

		fmt.Println(message)

		var cmd *Command
		var userName string
		var found bool
		errm := json.Unmarshal([]byte(message), &cmd)
		if errm != nil {

		} else {
			if cmd.Type == REGISTER {
				//check if user name already is taken.
				for d := range server.registeredConnections {
					if d.Usr != nil && d.Usr.Name == string(cmd.Value) {
						found = true
					}
				}
				if found {
					break
				} else {
					c.Usr = &User{Name: string(cmd.Value)}
				}
			}

			if cmd.Type == MESSAGE {
				//check if the message is a private message. 
				r, _ := regexp.MatchString("@", string(cmd.Value))
				if r {
					found = false
					userName = strings.Split(string(cmd.Value), " ")[0]
					userName = strings.Split(userName, "@")[1]
					for d := range server.registeredConnections {
						if d.Usr.Name == userName {
							d.send <- &Message{c.Usr, cmd.Value}
							found = true
						}
					}
					if found {
						c.send <- &Message{c.Usr, cmd.Value}
					} else {
						c.send <- &Message{&User{"From server"}, []byte("user not found")}
					}
					//else send the message to everyone 
				} else {
					server.broadcast <- &Message{c.Usr, cmd.Value}
				}
			}

			if cmd.Type == BLOGIN {
				server.broadcast <- &Message{c.Usr, []byte("Logged In")}
			}

			if cmd.Type == BLOGOUT {
				server.broadcast <- &Message{c.Usr, []byte("Logged Out")}
				break
			}

		}

	}
	c.Conn.Close()
}

func (c *Connection) Write() {
	for message := range c.send {
		jsonM, _ := json.Marshal(message)
		fmt.Println(string(jsonM))
		err := websocket.Message.Send(c.Conn, string(jsonM))
		if err != nil {
			break
		}
	}
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
