package users

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	helpers "github.com/M1n007/go-rest-api/repository/users/utils"
)

// ReturnAllUsers function
func ReturnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arrUser []Users
	var response ResponseGetUsers

	rows, err := GetAllUsers()

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal((err.Error()))
		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var response ResponseCreateUsers

	stmt, err := AddUsers()

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

	//converting keyVal map to jsonString
	newData := helpers.JsonStringConvert(keyVal)

	response.Status = 200
	response.Message = "Success Add Users!"
	response.Data = newData

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
