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

var index *template.Template

func init() {
	tdir := os.Getenv("TEMPLATE")
	log.Printf("Template Directory: %s\n", tdir)

	indexPath := fmt.Sprintf("%s%s", tdir, "index.html")
	index = template.Must(template.ParseFiles(indexPath))

}

const GofireSession = "gSession"

func checkSession(r *http.Request)string{
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
	w.Header().Set("content-type", "text/html")
	index.Execute(w, nil)
}
