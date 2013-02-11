package main

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"text/template"
	"os"
	"os/signal"
	//"strconv"
)

const addr = ":8080"

type User struct {
	Name string
}

type CommandType int

const (
	REGISTER = iota //0
	MESSAGE         //1
	BLOGIN
	BLOGOUT
)

type Command struct {
	Type  int
	Value []byte
}

type Connection struct {
	//User who sends
	Usr *User
	//websocket-connection
	Conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan *Message
}

type Message struct {
	Usr *User
	Msg []byte
}

type Server struct {
	history               []*Message
	broadcast             chan *Message
	register              chan *Connection
	unregister            chan *Connection
	registeredConnections map[*Connection]bool
}

const tempDir = "temp/"

func(s *Server) initDir(){
	err := os.Mkdir(tempDir, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func(s *Server) cleanUp(){
	err := os.RemoveAll(tempDir)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) run() {
	s.initDir()
	
	for {
		select {
		case c := <-server.register:
			server.registeredConnections[c] = true
			c.send <- &Message{&User{"From server"}, []byte("with love")}
			jsonU, _ := json.Marshal(s.history)
			//send history to new user 
			c.send <- &Message{&User{"All user"}, jsonU}
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

func (c *Connection) Read() {
	for {
		var message string
		err := websocket.Message.Receive(c.Conn, &message)
		if err != nil {
			break
		}

		fmt.Println(message)

		var cmd *Command
		var user_name string
		var found bool
		errm := json.Unmarshal([]byte(message), &cmd)
		if errm != nil {

		} else {
			if cmd.Type == REGISTER {
				//check if user name already is taken.
				for d := range server.registeredConnections {
					if d.Usr != nil && d.Usr.Name == string(cmd.Value) {
						found = true
					}
				}
				if found {
					break
				} else {
					c.Usr = &User{Name: string(cmd.Value)}
				}
			}

			if cmd.Type == MESSAGE {
				//check if the message is a private message. 
				r, _ := regexp.MatchString("@", string(cmd.Value))
				if r {
					found = false
					user_name = strings.Split(string(cmd.Value), " ")[0]
					user_name = strings.Split(user_name, "@")[1]
					for d := range server.registeredConnections {
						if d.Usr.Name == user_name {
							d.send <- &Message{c.Usr, cmd.Value}
							found = true
						}
					}
					if found {
						c.send <- &Message{c.Usr, cmd.Value}
					} else {
						c.send <- &Message{&User{"From server"}, []byte("user not found")}
					}
					//else send the message to everyone 
				} else {
					server.broadcast <- &Message{c.Usr, cmd.Value}
				}
			}

			if cmd.Type == BLOGIN {
				server.broadcast <- &Message{c.Usr, []byte("Logged In")}
			}

			if cmd.Type == BLOGOUT {
				server.broadcast <- &Message{c.Usr, []byte("Logged Out")}
				break
			}

		}

	}
	c.Conn.Close()
}

func (c *Connection) Write() {
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
	history:               make([]*Message, 0),
	broadcast:             make(chan *Message),
	register:              make(chan *Connection),
	unregister:            make(chan *Connection),
	registeredConnections: make(map[*Connection]bool),
}

func wsHandler(ws *websocket.Conn) {
	c := &Connection{Usr: nil, send: make(chan *Message), Conn: ws}
	server.register <- c
	defer func() { server.unregister <- c }()
	go c.Write()
	c.Read()
}

//var indexTemplate = template.Must(template.ParseFiles("template/index.html"))

const staticDir = "template/"

func mainHandler(w http.ResponseWriter, r *http.Request) {
	path := staticDir + r.URL.Path[1:]

	data, err := openFile(path)
	if err == nil {
		fmt.Println(path[len(path)-3:])
		if path[len(path)-3:] == "css" {

			w.Header().Set("Content-Type", STYLECSS)
		}

		if path[len(path)-2:] == "js" {

			w.Header().Set("Content-Type", JAVASCRIPT)
		}

		w.Write(data)
	} else {
		//using standard notfound impl.
		http.NotFound(w, r)
	}
}

const (
	TEXTPLAIN  string = "text/plain"
	STYLECSS   string = "text/css"
	JAVASCRIPT string = "application/x-javascript"
)

func openFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate, _ := template.ParseFiles("template/chat.html")
	indexTemplate.Execute(w, r.Host)
}

func doLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/chat.html", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile("template/index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(html)
}

//writes the file
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("File")

	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(tempDir+handler.Filename, data, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	go server.run()
	
	//cleanup if command+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		<-c
	    server.cleanUp()
		os.Exit(0)
	}()
	
	http.HandleFunc("/upload", uploadHandler)
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
