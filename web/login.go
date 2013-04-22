package web

import(
	"net/http"
	db "gofire/database"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.FormValue("user")
	password := r.FormValue("pw")

	token := db.IsUserPasswordValid(username, password)

	if token == ""{
		http.Redirect(w, r, "/", http.StatusTextOk)
		return
	}

	d := time.Now.Add(356*24*time.Hour)
	cookie := &http.Cookie{Name:GofireSession, Value:token, Expires:d, HttpOnly:true}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/chat", http.StatusFound)
	return
}
