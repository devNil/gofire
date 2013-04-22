package database

import(
	//"database/sql"
	//_ "github.com/bmizerany/pq"
	"log"
)

const qIsSessionValid = "select count(*) from gf_user where session=$1"

func IsSessionValid(token string)bool{
	conn := Open()
	defer conn.Close()

	var count int64

	row := conn.QueryRow(qIsSessionValid, token)
	
	err := row.Scan(&count)

	if err != nil{
		log.Println(err)
		return false
	}

	return count == 1

}

const qIsUserPasswordValid = "select session from gf_user where login=$1 and pw=$2 "

//Checks if user and password combination are valid
//returns "" if not, else the session token for the user
func IsUserPasswordValid(username, password string)string{
	conn := Open()
	defer conn.Close()

	var session string
	
	row := conn.QueryRow(qIsUserPasswordValid, username, password)

	err := row.Scan(&session)

	if err != nil{
		log.Println(err)
		return ""
	}

	return session
}
