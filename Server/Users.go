package Server

import (
	"GolangApi/DataAcess"
	"GolangApi/Model"
)

type User struct {
	Account  string
	Password string
	Email    string
}

var er = new(ErrorHandle)

func (u *User) CreateUser() {

	conn := DataAcess.Db

	stmt, err := conn.Prepare("Insert Users set account=? ,password = ? ,email = ?")
	er.CheckErr(err)

	stmt.Exec(u.Account, u.Password, u.Email)
	er.CheckErr(err)

}

func UsersList() []Model.Users {
	u := make([]Model.Users, 0)
	conn := DataAcess.Db

	row, err := conn.Query("SELECT * FROM users")
	er.CheckErr(err)

	for row.Next() {
		var userID int
		var account string
		var password string
		var email string

		row.Scan(&userID, &account, &password, &email)

		u = append(u, Model.Users{userID, account, password, email})
	}

	return u
}