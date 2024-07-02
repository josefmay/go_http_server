package main

import (
	"fmt"
	"os"
	"github.com/josefmay/go_http_server/cmd/api"
)


func main() {
	fmt.Printf("Hi")
	srv := api.initServer()
	err := srv.ListenAndServe()

	if err != nil{
		fmt.Printf("Error")
		os.Exit(1)
	}
}