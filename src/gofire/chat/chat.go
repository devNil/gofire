//This package provides all chatroom-functionality
package chat

import(
	"gofire/user"
)

type ChatRoom struct{
	Name string //The name of the chatroom, also the id of a chatroom
	User *[]user.User //List of all user in this chatroom
}
