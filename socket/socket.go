package socket

import(
	"gofire/web"
	db "gofire/database"
	"code.google.com/p/go.net/websocket"
	"log"
)

//good ol' days
type fireserver struct{
	connections map[*connection]bool
	broadcast chan message
	register chan *connection
	unregister chan *connection
}

var fs = &fireserver{
	make(map[*connection]bool),
	make(chan message),
	make(chan *connection),
	make(chan *connection),
}

func Start(){
	go fs.run()
}

func (f *fireserver)run(){
	for{
		select{
		case c := <-f.register:
			f.connections[c] = true
			c.send<-message{"Server", "Hello"}
		case c:= <-f.unregister:
			delete(f.connections, c)
			close(c.send)
		case m := <-f.broadcast:
			for c:= range f.connections{
				c.send<-m
			}
	}
}
}

type message struct{
	Username string
	Text string
}

type connection struct{
	user string
	conn *websocket.Conn
	send chan message
}

func (c *connection) sender(){
	for m := range c.send{
		err := websocket.JSON.Send(c.conn,m)
		if err != nil{
			log.Println(err)
			break
		}
	}
	c.conn.Close()
}

func (c *connection)reader(){
	for{
		var m message
		err := websocket.JSON.Receive(c.conn, &m)
		if err != nil{
			log.Println(err)
			break
		}

		m.Username = c.user

		fs.broadcast<-m
	}
	c.conn.Close()
}


func SocketHandler(conn *websocket.Conn){
	token := web.CheckSession(conn.Request())

	if token == ""{
		return
	}

	username := db.GetUser(token)

	if username == ""{
		return
	}

	c := &connection{
		username,
		conn,
		make(chan message),
	}

	fs.register<-c
	defer func(){fs.unregister<-c}()
	go c.sender()
	c.reader()
}
