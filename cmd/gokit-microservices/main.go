package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/katoozi/gokit-microservices/services/users"
)

func main() {
	router := mux.NewRouter()

	getUser := users.GetUserHandler()
	updateUser := users.UpdateUserHandler()
	deleteUser := users.DeleteUserHandler()
	createUser := users.CreateUserHandler()

	router.Handle("/user/{id}", getUser).Methods("GET")
	router.Handle("/user/{id}", updateUser).Methods("PUT")
	router.Handle("/user/{id}", deleteUser).Methods("DELETE")
	router.Handle("/user", createUser).Methods("CREATE")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
