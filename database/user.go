package database

import(
	"log"
)

const qGetUser = "select login from gf_user where session = $1"

func GetUser(token string)string{
	conn :=Open()
	defer conn.Close()

	var user string

	row := conn.QueryRow(qGetUser, token)

	err := row.Scan(&user)

	if err != nil{
		log.Println(err)
		return ""
	}

	return user


}
