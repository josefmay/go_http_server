package main

import (
	"log"

	"github.com/josefmay/go_http_server/cmd"
)


func main() {
	log.Print("Initializing server...")
	srv := api.InitServer()

	log.Print("Server starting....")
	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal("Server failed to launch.")
	}
}