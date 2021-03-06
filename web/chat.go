package web

import(
	"net/http"
)

func ChatHandler(w http.ResponseWriter, r *http.Request){
	token := CheckSession(r)

	if token == ""{
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	w.Header().Set("content-type", "text/html")
    templates.ExecuteTemplate(w,"chat", r.Host)
	return
}
