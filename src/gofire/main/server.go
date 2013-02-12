package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
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

	for {
		select {
		case c := <-server.register:
			server.registeredConnections[c] = true
			c.send <- &Message{&User{"From server"}, []byte("with love")}
			jsonU, _ := json.Marshal(s.history)
			//send history to new user 
			c.send <- &Message{&User{"history"}, jsonU}
		case c := <-server.unregister:
			delete(server.registeredConnections, c)
			close(c.send)
		case m := <-server.broadcast:
			//append to history
			s.history = append(s.history, m)
			for c := range s.registeredConnections {
				select {
				case c.send <- m:
				default:
					delete(s.registeredConnections, c)
					close(c.send)
					go c.Conn.Close()
				}
			}
		}
	}
}
