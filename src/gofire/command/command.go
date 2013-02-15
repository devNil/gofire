// Package command provides commands for querying both, server and client.
package command

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
