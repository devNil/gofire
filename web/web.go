//This packages consists all handler funcs for the webserver
package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	db "gofire/database"
)

var templates *template.Template

func init() {
	tdir := os.Getenv("TEMPLATE")
	log.Printf("Template Directory: %s\n", tdir)

    templates = template.Must(template.ParseGlob(tdir))

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
	index.Execute(w, nil)
}
