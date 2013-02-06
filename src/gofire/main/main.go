package main

import(
	"code.google.com/p/go.net/websocket"
	"net/http"
	"text/template"
	"encoding/json"
	"io/ioutil"
	"fmt"
	//"strconv"
)

const addr = ":8080"

type User struct{
	Name string
}

type CommandType int

const(
	REGISTER = iota //0
	MESSAGE			//1
	GETHISTORY		//2
	BLOGIN
	BLOGOUT
)

type Command struct{
	Type int
	Value []byte
}

type Connection struct{
	//User who sends
	Usr *User
	//websocket-connection
	Conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan *Message
}

type Message struct{
	Usr *User
	Msg string
}

type Server struct{
	history []*Message
	broadcast chan *Message
	register chan *Connection
	unregister chan *Connection
	registeredConnections map[*Connection]bool
}

func (s *Server)run(){
	for {
		select {
		case c := <-server.register:
			server.registeredConnections[c] = true
			c.send<-&Message{&User{"From server"}, "with love"}
		case c := <-server.unregister:
			delete(server.registeredConnections, c)
			close(c.send)
		case m := <-server.broadcast:
			//append to history
			s.history = append(s.history, m)
			for c := range s.registeredConnections {
				select {
				case c.send <- m:
				default:
					delete(s.registeredConnections, c)
					close(c.send)
					go c.Conn.Close()
				}
			}
		}
	}
}

func (c *Connection)Read(){
	for {
		var message string
		err := websocket.Message.Receive(c.Conn, &message)
		if err != nil {
			break
		}

		fmt.Println(message)

		var cmd *Command

		errm := json.Unmarshal([]byte(message), &cmd)
		if(errm != nil){

		}else{
			if cmd.Type == REGISTER {
				c.Usr = &User{Name:string(cmd.Value)}
			}

			if cmd.Type == MESSAGE {
				server.broadcast<-&Message{c.Usr, string(cmd.Value)}
			}

			if cmd.Type == GETHISTORY {
				r, _ := json.Marshal(server.history)
				c.send<-&Message{&User{"history"}, string(r)}
			}
			if cmd.Type == BLOGIN {
				server.broadcast<-&Message{c.Usr, string("Logged In")}
			}

			if cmd.Type == BLOGOUT {
				fmt.Println("Somebody wants to logout")
			}

			//c.Usr = user
		}

	}
	c.Conn.Close()
}

func (c *Connection)Write(){
	for message := range c.send {
		jsonM, _ := json.Marshal(message)
		fmt.Println(string(jsonM))
		err := websocket.Message.Send(c.Conn, string(jsonM))
		if err != nil {
			break
		}
	}
}

var server = Server{
	history: make([]*Message, 0),
	broadcast: make(chan *Message),
	register: make(chan *Connection),
	unregister: make(chan *Connection),
	registeredConnections: make(map[*Connection]bool),
}

func wsHandler(ws *websocket.Conn) {
	c := &Connection{Usr: nil,send: make(chan *Message), Conn: ws}
	server.register <- c
	defer func() {server.unregister <- c}()
	go c.Write()
	c.Read()
}

//var indexTemplate = template.Must(template.ParseFiles("template/index.html"))

func mainHandler(w http.ResponseWriter, r *http.Request) {

}

func chatHandler(w http.ResponseWriter, r *http.Request){
	indexTemplate, _ := template.ParseFiles("template/chat.html")
	indexTemplate.Execute(w, r.Host)
}

func doLogin(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	pendingUser = &User{r.Form["username"][0]}

	http.Redirect(w, r, "/chat.html", http.StatusFound)
}

var pendingUser *User

func loginHandler(w http.ResponseWriter, r *http.Request){
	html, err := ioutil.ReadFile("template/index.html")
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(html)
}

func main(){
	go server.run()
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/chat.html", chatHandler)
	http.HandleFunc("/doLogin", doLogin)
	http.Handle("/ws", websocket.Handler(wsHandler))
	err := http.ListenAndServe(addr, nil)
	    if err != nil {
	        panic("ListenAndServe: " + err.Error())
	    }
}
