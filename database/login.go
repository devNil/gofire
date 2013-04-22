package database

import(
	//"database/sql"
	//_ "github.com/bmizerany/pq"
	"log"
)

const qIsSessionValid = "select count(*) from gf_users where session=$1"

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
