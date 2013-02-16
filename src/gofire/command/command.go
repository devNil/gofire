// Package command provides commands for querying both, server and client.
package command

import (
	"gofire/message"
	"gofire/user"
	"encoding/json"
)

type CommandType int

const (
	REGISTER CommandType = iota //Register a user
	//Text-Messaging Commands
	BMESSAGE //Broadcast-Message
	UMESSAGE //Unicast-Message
	MMESSAGE //Multicast-Message
	//System Commands
	BLOGIN  //Broadcast-Login
	BLOGOUT //Broadcast-Logout
)

//Every Command has a type and a value for sending arbitrary data.
type Command struct {
	Type  CommandType
	Value []byte
}

//This Function prepares wraps a message in an command.
//This is a shortcut.
//If something went wrong, nil and a error are returned.
func PrepareMessage(tp CommandType ,usr *user.User, msg []byte) (*Command,error){
	mMessage, err := json.Marshal(message.Message{User:usr,Msg:msg})
	if err != nil{
		return nil, err
	}
	
	return &Command{tp, mMessage}, nil
}
