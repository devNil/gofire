package main

import (
	"encoding/json"
	"fmt"
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
	history               []*Message
	broadcast             chan *Message
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
	for _,cr := range s.chatRooms {
		go cr.run()
	}
	for {
		select {
		case c := <-s.register:
			s.registeredConnections[c] = true
			s.chatRooms[0].register <- c
			c.send <- &Message{&User{"From server"}, []byte("with love")}
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
				c.send <- &Message{&User{cr.name}, []byte("go!")}
				jsonU, _ := json.Marshal(cr.history)
				//send history to new user 
				c.send <- &Message{&User{"history"}, jsonU}
			case c := <-cr.unregister:
				delete(cr.registeredConnections, c)
			
			case m:= <-cr.broadcast:
				//append to history
				cr.history = append(cr.history, m)
				for c := range cr.registeredConnections {
					select {
						case c.send <- m:
							
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
		history:               make([]*Message, 0),
		broadcast:             make(chan *Message),
		register:              make(chan *Connection),
		unregister:            make(chan *Connection),
		registeredConnections: make(map[*Connection]bool),
	})	
}
