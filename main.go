package main

import (
	"fmt"

	server "github.com/M1n007/go-rest-api/server"
)

func main() {
	server.Init()
	fmt.Println("Server is running on port 9000")
}
