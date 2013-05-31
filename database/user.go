package database

//User object retrieved from database
type User struct{
    Id int64
    Login string
    Admin bool
    Session string
}

const qGetUser = "select id, login, mod, session from gf_user where id = $1"

func GetUser(id int64)(*User, error){
    conn := Open()
    defer conn.Close()
    
    user = new(User)
    var mod int64

    row := conn.QueryRow(qGetUser, id)
    err := row.Scan(&user.Id, &user.Login, &mod, &user.Session)

    &user.Admin = mod == 1

    return &user
}

}
