package main

import (
	"encoding/json"
	"fmt"
	"gofire/command"
	"gofire/message"
	"gofire/user"
	//"regexp"
	//"strings"
	"code.google.com/p/go.net/websocket"
)

type Connection struct {
	//User who sends
	Usr *user.User
	//websocket-connection
	Conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan *command.Command
	// chatroom in which the user is
	chatRoom *ChatRoom
}

func (c *Connection) Read() {
	for {
		//raw incoming command
		var rawIncome []byte

		err := websocket.Message.Receive(c.Conn, &rawIncome)
		if err != nil {
			fmt.Println(err)
			break
		}

		var cmd *command.Command

		//var userName string
		var found bool

		err = json.Unmarshal([]byte(rawIncome), &cmd)

		if err == nil {
			//REGISTER -> {REGISTER, USERNAME}
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
					c.Usr = &user.User{Name: string(cmd.Value)}
				}
			}

			if cmd.Type == command.BMESSAGE {
				m, errm := json.Marshal(message.Message{c.Usr, cmd.Value})
				if errm == nil {
					c.chatRoom.broadcast <- &command.Command{command.BMESSAGE, m}
				} else {
					fmt.Println(errm)
				}
			}

			/*if cmd.Type == command.MESSAGE {
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
			}*/
		}

		if cmd.Type == command.BLOGIN {
			fmt.Println(c.Usr)
			cmd,err := command.PrepareMessage(command.BMESSAGE, c.Usr, []byte("Logged In"))
			if err == nil {
				c.chatRoom.broadcast <- cmd
			} else {
				fmt.Println(err)
			}

		}
		
		if cmd.Type == command.BLOGOUT {
			cmd,err := command.PrepareMessage(command.BMESSAGE, c.Usr, []byte("Logged out"))
			if err == nil {
				c.chatRoom.broadcast <- cmd
			} else {
				fmt.Println(err)
			}

		}
	}

	c.Conn.Close()
}

func (c *Connection) Write() {
	for command := range c.send {
		//marshal in 

		jsonC, _ := json.Marshal(command)

		err := websocket.Message.Send(c.Conn, string(jsonC))
		if err != nil {
			break
		}
	}
}
