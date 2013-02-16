//Package message provides the message structs
package message

import "gofire/user"

type Message struct {
	User *user.User //User who sended the message
	Msg  []byte     //For arbitrary data
}
