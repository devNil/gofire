//This package contains all the stuff related to the database
//it abstracts most of the database related tasks like login of a user
package database

import (
	sha "crypto/sha512"
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
	"log"
	"os"
)

//the connection url for the database driver
//the momentary used database system is postgres because its cool
const url = "user=%s password=%s dbname=%s sslmode=%s host=%s"

var dUser, dDb, dPw, dHost string

var dMode = "disable"

func init() {
	dUser = os.Getenv("DUSER")
	dDb = os.Getenv("DDB")
	dPw = os.Getenv("DPW")
	dHost = os.Getenv("DHOST")

	if mode := os.Getenv("DMODE"); mode != "" {
		dMode = mode
	}

	log.Printf("Connection to database with: %s\n", fmt.Sprintf(url, dUser, dPw, dDb, dMode, dHost))
}

//helper function to generate sha512 hex-string-hashes
func sha512(input string) string {
	h := sha.New()
	r := h.Sum([]byte(input))
	return fmt.Sprintf("%x", r)
}

//Opens a connection to the database
//if an error occurs, gofire panics
func Open() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf(url, dUser, dPw, dDb, dMode, dHost))
	if err != nil{
		panic(err)
	}

	return db
}
