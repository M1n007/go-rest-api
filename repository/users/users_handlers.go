package users

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	DB "github.com/M1n007/go-rest-api/helpers/database/mysql"
)

// ReturnAllUsers function
func ReturnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response Response

	db := DB.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM person")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal((err.Error()))
		} else {
			arr_user = append(arr_user, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var users Users
	var arr_user []Users
	var response Response

	db := DB.Connect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO person(id,first_name,last_name) VALUES(?,?,?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	// change body to key value
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	// identifier key value
	id := keyVal["id"]
	firstname := keyVal["firstname"]
	lastname := keyVal["lastname"]

	// execution stmt for inserting data to person
	_, err = stmt.Exec(id, firstname, lastname)

	if err != nil {
		log.Fatal(err.Error())
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	json.NewEncoder(w).Encode(response)
}
