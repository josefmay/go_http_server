package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
/account authentication JWT
/watchlist watchlist database
/python backtester optimization engine
/report run reporting

sqlite3

react frontend?
*/

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
		log.Println(r.URL.Path)
        fmt.Fprintf(w, "Welcome to my website!")
    })

	http.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	http.ListenAndServe(s.listenAddr, nil)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	log.Println(r.URL.Path)

	//Switch statement later

	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)

	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)

	}

	return fmt.Errorf("Not a valid method")
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	ToJSON(w, 200, "Handled like a boss")
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToJSON(w http.ResponseWriter, status int, v any) error {
	fmt.Println("Converting to msg JSON...")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}


type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			ToJSON(w, http.StatusBadRequest, err.Error())
		}
	}
}
