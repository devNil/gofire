package web

import(
	"net/http"
)

func ChatHandler(w http.ResponseWriter, r *http.Request){

    session, err := store.Get(r, cookieName)

    if err != nil || session.IsNew {
        http.Redirect(w,r, "/", http.StatusFound)
        return
    }

	w.Header().Set("content-type", "text/html")
    templates.ExecuteTemplate(w,"chat", r.Host)
	return
}
