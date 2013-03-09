//This package provides all chatroom-functionality
package chat

import(
	"gofire/user"
	"gofire/message"
	"code.google.com/p/go.net/websocket"
)

type Connection struct{
	User *user.User //User in connection
	conn *websocket.Conn
	send chan *command.Command
}

type ChatRoom struct{
	Name string //The name of the chatroom, also the id of a chatroom
	register chan *Connection
	unregister chan *Connection
	registeredConnections map[*Connection]bool
}
