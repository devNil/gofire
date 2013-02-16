package main

import (
	"fmt"
	"gofire/command"
	"gofire/message"
	"gofire/user"
	"os"
)

type Server struct {
	chatRooms             []*ChatRoom
	register              chan *Connection
	unregister            chan *Connection
	registeredConnections map[*Connection]bool
}

type ChatRoom struct {
	name                  string
	history               []*message.Message
	broadcast             chan *command.Command
	register              chan *Connection
	unregister            chan *Connection
	registeredConnections map[*Connection]bool
}

func (s *Server) initDir() {
	err := os.Mkdir(TEMPDIR, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) cleanUp() {
	err := os.RemoveAll(TEMPDIR)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) run() {
	s.initDir()
	s.chatRooms = append(s.chatRooms, &chatRoom)
	for _, cr := range s.chatRooms {
		go cr.run()
	}
	for {
		select {
		case c := <-s.register:
			s.registeredConnections[c] = true
			s.chatRooms[0].register <- c
			command, err := command.PrepareMessage(command.BMESSAGE, &user.User{"From server"}, []byte("with love"))

			if err == nil {
				c.send <- command
			} else {
				fmt.Println(err)
			}

		case c := <-s.unregister:
			s.chatRooms[0].unregister <- c
			delete(s.registeredConnections, c)
			close(c.send)
		}
	}
}

func (cr *ChatRoom) run() {
	for {
		select {
		case c := <-cr.register:
			cr.registeredConnections[c] = true

			cmd, err := command.PrepareMessage(command.BMESSAGE, &user.User{cr.name}, []byte("go!"))

			if err == nil {
				c.send <- cmd
			} else {
				fmt.Println(err)
			}

			//jsonU, _ := json.Marshal(cr.history)
			//send history to new user 
			//c.send <- &Message{&User{"history"}, jsonU}
		case c := <-cr.unregister:
			delete(cr.registeredConnections, c)

		case cmd := <-cr.broadcast:
			//append to history
			//cr.history = append(cr.history, m)
			for c := range cr.registeredConnections {
				select {
				case c.send <- cmd:

				default:
					delete(cr.registeredConnections, c)
					server.unregister <- c
					go c.Conn.Close()
				}
			}
		}
	}
}
func (s *Server) creatChatRoom(name string) {
	s.chatRooms = append(s.chatRooms, &ChatRoom{
		name:                  name,
		history:               make([]*message.Message, 0),
		broadcast:             make(chan *command.Command),
		register:              make(chan *Connection),
		unregister:            make(chan *Connection),
		registeredConnections: make(map[*Connection]bool),
	})
}
