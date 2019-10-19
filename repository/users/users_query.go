package users

import (
	"database/sql"

	DB "github.com/M1n007/go-rest-api/helpers/database/mysql"
)

// GetAllUsers ...
func GetAllUsers() (rows *sql.Rows, err error) {

	db := DB.Connect()

	defer db.Close()

	rows, err = db.Query("SELECT * FROM person")

	return rows, err
}

// AddUsers ...
func AddUsers() (stmt *sql.Stmt, err error) {
	db := DB.Connect()

	stmt, err = db.Prepare("INSERT INTO person(id,first_name,last_name) VALUES(?,?,?)")

	return stmt, err
}
