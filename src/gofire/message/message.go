//will all message related stuff
package message

import "gofire/user"

type Message struct {
	User *user.User //User who sended the message
	Msg []byte //For arbitrary data
}