package web

import(
	"net/http"
    "gofire/database"
    "log"
)

func ChatHandler(w http.ResponseWriter, r *http.Request){

    session, err := store.Get(r, cookieName)

    // Pretty useless atm. Gorilla restores the sessions
	// values from the clients cookie value
	if session.Values["id"] == nil || session.IsNew {
		session.Options.MaxAge = -1
		session.Save(r, w)
        http.Redirect(w,r, "/", http.StatusFound)
        return
    }

    id, _ := session.Values["id"].(int64)

    user , err := database.GetUser(id)

    if err != nil{
        log.Println(err)
        http.Redirect(w,r, "/", http.StatusFound)
        return
    }

	w.Header().Set("content-type", "text/html")
    values := map[interface{}]interface{}{
        "Host":r.Host,
        "User":user,
    }

    err = templates.ExecuteTemplate(w,"chat", values)
    log.Println(err)
	return
}
