package web

import (
	"net/http"
)
func LogoutHandler(w http.ResponseWriter, r *http.Request){

    session, _ := store.Get(r, cookieName)

    session.Options.MaxAge = -1

    store.Save(r, w, session)

	http.Redirect(w, r, "/", http.StatusFound)
}
