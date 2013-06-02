package web

import(
	"net/http"
    "log"
    "gofire/database"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.FormValue("user")
	password := r.FormValue("pw")

    id, err := database.GetUserId(username, password)

    log.Println(id)

    if err != nil{
        log.Println(err)
        http.Redirect(w, r, "/", http.StatusFound)
        return
    }

    session , err := store.Get(r, cookieName)

    if err != nil{
        log.Println(err)
        http.Redirect(w, r, "/", http.StatusFound)
        return
    }

    session.Values["id"] = id

    store.Save(r, w, session)

	http.Redirect(w, r, "/chat", http.StatusFound)
	return
}
