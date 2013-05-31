package web

import(
	"net/http"
	"time"
    "log"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.FormValue("user")
	password := r.FormValue("pw")

	//token := db.IsUserPasswordValid(username, password)

	/*if token == ""{
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}*/

	http.Redirect(w, r, "/chat", http.StatusFound)
	return
}
