const REGISTER = 0 //Register a user
//Text-Messaging Commands
const BMESSAGE = 1 //Broadcast-Message
const UMESSAGE = 2 //Unicast-Message
const MMESSAGE = 3 //Multicast-Message
	
const BLOGIN = 4 //Broadcast-Login
const BLOGOUT = 5 //Broadcast-Logout

function Command(type, value){
	this.type = type;
	this.value = value;
}

Command.prototype.Send = function(websocket){
	this.value = window.btoa(this.value);
	websocket.send(JSON.stringify(this));
}