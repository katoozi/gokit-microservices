package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/katoozi/gokit-microservices/services/auth"
	"github.com/katoozi/gokit-microservices/services/users"
)

func main() {
	router := mux.NewRouter()

	userUrls := router.PathPrefix("/user").Subrouter()
	userUrls.Handle("/", users.CreateUserHandler()).Methods("CREATE")
	userUrls.Handle("/{id}", users.GetUserHandler()).Methods("GET")
	userUrls.Handle("/{id}", users.UpdateUserHandler()).Methods("PUT")
	userUrls.Handle("/{id}", users.DeleteUserHandler()).Methods("DELETE")
	userUrls.Use(auth.HeaderValidation)

	secret := "askdmajklsdnjkansdjknansdkajsndkjnaskjdnkajsnd"
	authUrls := router.PathPrefix("/auth").Subrouter()
	authUrls.Handle("/login", auth.LoginHandler(secret)).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
