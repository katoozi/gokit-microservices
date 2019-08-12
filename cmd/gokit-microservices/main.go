package main

import (
	"log"
	"net/http"

	"github.com/katoozi/gokit-microservices/services/users"
)

func main() {
	getUser := users.GetUserHandler()
	updateUser := users.UpdateUserHandler()
	deleteUser := users.DeleteUserHandler()

	http.Handle("/user/get", getUser)
	http.Handle("/user/update", updateUser)
	http.Handle("/user/delete", deleteUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
