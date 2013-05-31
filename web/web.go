//This packages consists all handler funcs for the webserver
package web

import (
	"html/template"
	"log"
	"net/http"
	"os"
	db "gofire/database"
    "fmt"
    "github.com/gorilla/sessions"
)

var templates *template.Template

var staticDir string

const cookieName = "this-is-the-most-awesome-cookie-name"

//memory-cookiename
var store = sessions.NewCookieStore([]byte("you-cannot-hack-this"))

func init() {
	tdir := os.Getenv("TEMPLATE")
	log.Printf("Template Directory: %s\n", tdir)

    templates = template.Must(template.ParseGlob(tdir))

    staticDir = os.Getenv("STATIC")
    log.Println("Static Dir: ", staticDir)
}

const GofireSession = "gSession"

func CheckSession(r *http.Request)string{
	if cookie, err :=r.Cookie(GofireSession); err != nil{
		return ""
	}else{
		if db.IsSessionValid(cookie.Value){
			return cookie.Value
		}
		return ""
	}
	return ""
}


//Handler for the index-site
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//Session validation
	if token := CheckSession(r); token != ""{
		http.Redirect(w, r, "/chat", http.StatusFound)
		return
	}
	w.Header().Set("content-type", "text/html")
    templates.ExecuteTemplate(w, "login", nil)
}

func StaticHandler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r, fmt.Sprint(staticDir, r.URL.Path[1:]))
}
