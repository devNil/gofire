package database

//User object retrieved from database
type User struct{
    Id int64
    Login string
    Admin bool
    Session string
}

const qGetUser = "select id, login, mod, session from gf_user where id = $1"

//retrieve userobject with id
func GetUser(id int64)(*User, error){
    conn := Open()
    defer conn.Close()
    
    user := new(User)
    var mod int64

    row := conn.QueryRow(qGetUser, id)
    err := row.Scan(&user.Id, &user.Login, &mod, &user.Session)

    user.Admin = mod == 1

    return user, err
}

const qGetUserId = "select id from gf_user where login=$1 and pw=$2"

//get user id with username and password combination
func GetUserId(username, password string)(int64, error){
    conn := Open()
    defer conn.Close()

    var id int64

    row := conn.QueryRow(qGetUserId, username, sha512(password))

    err := row.Scan(&id)

    return id, err
}
