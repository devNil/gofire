package main

import(
	"fmt"
	"log"
	"net/http"
	"os"
	"gofire/web"
)

const StandardPort = "8080"

func main(){

	port := os.Getenv("PORT")

	if port == ""{
		port = StandardPort
	}

	http.HandleFunc("/", web.IndexHandler)
	log.Println("IndexHandler registered")

	log.Printf("Server started on port: :%s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port),nil)
	panic(err)
}
