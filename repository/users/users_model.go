package users

//Users models
type Users struct {
	Id        string `form:"id" json:"id"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
}

// ResponseGetUsers ...
type ResponseGetUsers struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

// ResponseCreateUsers ...
type ResponseCreateUsers struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string
}
