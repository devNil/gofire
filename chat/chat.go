//This package provides all chatroom-functionality
package chat

import(
	"gofire/user"
	"gofire/message"
	"code.google.com/p/go.net/websocket"
)

type Connection struct{
	User *user.User //User in connection
	conn *websocket.Conn //a websocket-connection
	send chan *command.Command
}

type ChatRoom struct{
	Name string //The name of the chatroom, also the id of a chatroom
	register chan *Connection //Channel for incoming user
	unregister chan *Connection //Channel for "logout-user"
	registeredConnections map[*Connection]bool //Every Connection to a specific chatroom is tracked down here
}

func (chatroom *ChatRoom) Run(){
	for{
		select{
		case c <- chatroom.register:
			chatroom.registeredConnections[c] = true
			//TODO  code for broadcasting login 
			break
		case c <- chatroom.unregister:
			delete(chatroom.registeredConnections, c)
			//TODO code for broadcasting logout
			break
		}
	}
}
