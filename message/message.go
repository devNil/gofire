//Package message provides the message structs
package message

import "gofire/user"

type MessageType int

const(
	LOGIN MessageType = iota //User logged in message
)

type Message struct {
	User *user.User //User who sended the message
	Msg  []byte     //For arbitrary data
}
