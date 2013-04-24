package web

import (
	"net/http"
	"log"
)
func LogoutHandler(w http.ResponseWriter, r *http.Request){
    if cookie, err :=r.Cookie(GofireSession); err != nil{
		log.Println(err)
	} else {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	    http.Redirect(w, r, "/", http.StatusFound)
	}
}
