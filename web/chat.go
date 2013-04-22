package web

import(
	"net/http"
)

func ChatHandler(w http.ResponseWriter, r *http.Request){
	token := CheckSession(r)

	if token == ""{
		http.Redirect(w, r, "/", http.StatusOk)
		return
	}

	chat.Execute(w, r.Host)
}
