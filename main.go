package main

import (
	"fmt"

	server "github.com/M1n007/go-rest-api/server"
)

func main() {
	fmt.Printf("Server is running on port 9000")
	server.Init()
}
