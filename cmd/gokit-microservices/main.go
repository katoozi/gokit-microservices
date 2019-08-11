package main

import (
	"log"
	"net/http"

	"github.com/katoozi/gokit-microservices/services/users"
)

func main() {
	getUser := users.GetUserHandler()

	http.Handle("/user/get", getUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
