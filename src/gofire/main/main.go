package main

import(
	"code.google.com/p/go.net/websocket"
	"net/http"
	"text/template"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
	"regexp"
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
	Msg []byte
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
			c.send<-&Message{&User{"From server"}, []byte("with love")}
			
			jsonU, _ := json.Marshal(s.history)
			c.send<-&Message{&User{"All user"}, jsonU}
			//send the history to a new user
			/*for _, message := range s.history {
				c.send<- message
			}*/
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
		var user_name string
		errm := json.Unmarshal([]byte(message), &cmd)
		if(errm != nil){

		}else{
			if cmd.Type == REGISTER {
				c.Usr = &User{Name:string(cmd.Value)}
			}

			if cmd.Type == MESSAGE {
				r,_ := regexp.MatchString("@", string(cmd.Value))
				if r {
					var found bool = false
					user_name = strings.Split(string(cmd.Value), " ")[0]
					user_name = strings.Split(user_name, "@")[1]
					fmt.Println(user_name)
					for d := range server.registeredConnections {
						if d.Usr.Name == user_name {
							d.send <-&Message{c.Usr, cmd.Value}
							found = true
						}
					}
					if(found) {
						c.send <-&Message{c.Usr, cmd.Value}
					} else {
						c.send <-&Message{&User{"From server"},[]byte("user not found")}
					}
				} else {
					server.broadcast<-&Message{c.Usr, cmd.Value}
				}
			}

			if cmd.Type == BLOGIN {
				server.broadcast<-&Message{c.Usr, []byte("Logged In")}
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
	http.Redirect(w, r, "/chat.html", http.StatusFound)
}

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
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/index.html", loginHandler)
	http.HandleFunc("/chat.html", chatHandler)
	http.HandleFunc("/doLogin", doLogin)
	http.Handle("/ws", websocket.Handler(wsHandler))
	err := http.ListenAndServe(addr, nil)
	    if err != nil {
	        panic("ListenAndServe: " + err.Error())
	    }
}
