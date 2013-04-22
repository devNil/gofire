//This packages consists all handler funcs for the webserver
package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var index *template.Template

func init() {
	tdir := os.Getenv("TEMPLATE")
	log.Printf("Template Directory: %s\n", tdir)

	indexPath := fmt.Sprintf("%s%s", tdir, "index.html")
	index = template.Must(template.ParseFiles(indexPath))

}

//Handler for the index-site
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//Session validation
	w.Header().Set("content-type", "text/html")
	index.Execute(w, nil)
}
