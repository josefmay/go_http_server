package main

import (
	"fmt"
	
	"net/http"
)

type APIServer struct {
	listenAddr string
}

func newAPIServer(listenAddr string) * APIServer{
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

	http.ListenAndServe(s.listenAddr, nil)
}

