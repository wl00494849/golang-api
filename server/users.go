package server

import (
	"fmt"
	"golang-api/dataAccess"
	"golang-api/model"
)

type User struct {
	Account  string
	Password string
	Email    string
}

func (u *User) CreateUser() {

	conn := dataAccess.Db

	stmt, err := conn.Prepare("Insert Users set account=? ,password = ? ,email = ?")
	er.CheckErr(err)

	stmt.Exec(u.Account, u.Password, u.Email)

}

func DeleteUser(userID int) {
	conn := dataAccess.Db

	stmt, err := conn.Prepare("Delete from Users where UserID = ?")
	er.CheckErr(err)

	stmt.Exec(userID)

	fmt.Println("sucess")
}

func UsersList() []model.User {
	userSlice := make([]model.User, 0)
	conn := dataAccess.Db

	row, err := conn.Query("SELECT * FROM users")

	er.CheckErr(err)

	//func結束前關閉
	defer row.Close()
	for row.Next() {
		user := new(model.User)

		ScanToStruct(row, user)

		userSlice = append(userSlice, *user)
	}

	return userSlice
}
