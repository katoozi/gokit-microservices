package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/katoozi/gokit-microservices/services/users"
)

func main() {
	r := mux.NewRouter()

	getUser := users.GetUserHandler()
	updateUser := users.UpdateUserHandler()
	deleteUser := users.DeleteUserHandler()

	r.Handle("/user/{id}", getUser).Methods("GET")
	r.Handle("/user/{id}", updateUser).Methods("PUT")
	r.Handle("/user/{id}", deleteUser).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
