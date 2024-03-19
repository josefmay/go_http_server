package main

import (
	"fmt"
)

func main() {
	server := newAPIServer(":3000")
	fmt.Println("server is running on port ", server.listenAddr)
	server.Run()
}