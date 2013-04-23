package web

import (
	"net/http"
//	"time"
)
func LogoutHandler(w http.ResponseWriter, r *http.Request){
    if cookie, err :=r.Cookie(GofireSession); err != nil{
		return
	} else {
		cookie.MaxAge = -1
		//cookie.Expires = time.Now()
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}
