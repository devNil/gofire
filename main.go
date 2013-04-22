package main

import(
	"fmt"
	"log"
	"net/http"
	"os"
)

const StandardPort = "8080"

func main(){

	port := os.Getenv("PORT")

	if port == ""{
		port = StandardPort
	}

	log.Printf("Server started on port: :%s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port),nil)
	panic(err)
}
