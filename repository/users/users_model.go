package users

//Users models
type Users struct {
	Id        string `form:"id" json:"id"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
}

//Response models
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}
