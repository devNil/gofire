//This package provides basic server rest methods
package api

import(
	"net/http"
	"gofire/server"
)

func Ping (fireserver *server.FireServer,w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Ping-Api-Call"))
}
