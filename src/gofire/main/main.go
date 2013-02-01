package main

import(
	"code.google.com/p/go.net/websocket"
	"net/http"
	"text/template"
	"encoding/json"
	"io/ioutil"
	"crypto/sha1"
	"fmt"
)

type User struct{
	Name string
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
			c.send<-&Message{&User{"Tha server"}, "Welcome"}
		case c := <-server.unregister:
			delete(server.registeredConnections, c)
			close(c.send)
		case m := <-server.broadcast:
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
		//TODO user
		server.broadcast<-&Message{c.Usr, message}
	}
	c.Conn.Close()
}

func (c *Connection)Write(){
	for message := range c.send {
		jsonM, _ := json.Marshal(message)
		err := websocket.Message.Send(c.Conn, string(jsonM))
		if err != nil {
			break
		}
	}
}

var server = Server{
	broadcast: make(chan *Message),
	register: make(chan *Connection),
	unregister: make(chan *Connection),
	registeredConnections: make(map[*Connection]bool),
}

func wsHandler(ws *websocket.Conn) {
	c := &Connection{Usr: pendingUser,send: make(chan *Message), Conn: ws}
	server.register <- c
	defer func() {server.unregister <- c}()
	go c.Write()
	c.Read()
}

//var indexTemplate = template.Must(template.ParseFiles("template/index.html"))

func mainHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("template/index.html")
	indexTemplate.Execute(w, r.Host)
}

func doLogin(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	
	pendingUser = &User{r.Form["username"][0]}
	
	http.Redirect(w, r, "/index.html", http.StatusFound)
}

var pendingUser *User

func loginHandler(w http.ResponseWriter, r *http.Request){
	html, err := ioutil.ReadFile("template/login.html")
	if err != nil{
		w.Write([]byte("Fehler"))
		return
	}
	h := sha1.New()
	
	fmt.Println(h.Sum([]byte("Hello")))
	
	w.Write(html)
}

func main(){
	go server.run()
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/login.html", loginHandler)
	http.HandleFunc("/doLogin", doLogin)
	http.Handle("/ws", websocket.Handler(wsHandler))
	err := http.ListenAndServe(":8080", nil)
	    if err != nil {
	        panic("ListenAndServe: " + err.Error())
	    }
}