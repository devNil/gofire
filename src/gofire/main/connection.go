package main

import(
	"fmt"
	"gofire/command"
	"encoding/json"
	"regexp"
	"strings"
	"code.google.com/p/go.net/websocket"
)

type User struct {
	Name string
}

type Connection struct {
	//User who sends
	Usr *User
	//websocket-connection
	Conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan *Message
	// chatroom in which the user is
	chatRoom *ChatRoom
}

func (c *Connection) Read() {
	for {
		//var message string
		var rawIncome []byte

		err := websocket.Message.Receive(c.Conn, &rawIncome)
		if err != nil {
			fmt.Println(err)
			break
		}

		var cmd *command.Command
		var userName string
		var found bool
		errm := json.Unmarshal([]byte(rawIncome), &cmd)
		if errm != nil {

		} else {
			if cmd.Type == command.REGISTER {
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

			if cmd.Type == command.MESSAGE {
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
					c.chatRoom.broadcast <- &Message{c.Usr, cmd.Value}
				}
			}

			if cmd.Type == command.BLOGIN {
				c.chatRoom.broadcast <- &Message{c.Usr, []byte("Logged In")}
			}

			if cmd.Type == command.BLOGOUT {
				c.chatRoom.broadcast <- &Message{c.Usr, []byte("Logged Out")}
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
