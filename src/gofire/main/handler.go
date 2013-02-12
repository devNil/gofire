package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

const (
	TEXTPLAIN  string = "text/plain"
	STYLECSS   string = "text/css"
	JAVASCRIPT string = "application/x-javascript"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index.html", http.StatusFound)
		return
	}
	fmt.Println(r.URL.Path)
	path := STATICDIR + r.URL.Path[1:]

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

	err = ioutil.WriteFile(TEMPDIR+handler.Filename, data, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

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

func wsHandler(ws *websocket.Conn) {
	c := &Connection{Usr: nil, send: make(chan *Message), Conn: ws}
	server.register <- c
	defer func() { server.unregister <- c }()
	go c.Write()
	c.Read()
}
