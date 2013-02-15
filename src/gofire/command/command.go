// Package command provides commands for querying both, server and client.
package command

type CommandType int

const (
	REGISTER CommandType = iota //Register a user
	MESSAGE                     //Message in Value
	BLOGIN
	BLOGOUT
	BMESSAGE //Broadcast-Message
)

//Every Command has a type and a value for sending arbitrary data.
type Command struct {
	Type  CommandType
	Value []byte
}
