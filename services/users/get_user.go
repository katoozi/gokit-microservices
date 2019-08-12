package users

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// GetUserResponse is the GetUser reponse structure
type GetUserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// NewGetUserResponse is the GetUserResponse factory function
func NewGetUserResponse(id, username, firstName, lastName string, age int) GetUserResponse {
	return GetUserResponse{
		ID:        id,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

// GetUserReqeust is the GetUser request structure
type GetUserReqeust struct {
	ID string `json:"id"`
}

func (ser service) GetUser(user GetUserReqeust) (GetUserResponse, error) {
	if user.ID == "" {
		return GetUserResponse{}, errEmptyUsername
	}
	// TODO: check user does exsist or not(do sql query for example).
	// TODO: if user exists, also check the password.
	resp := NewGetUserResponse(
		user.ID,
		"download",
		"mohammad",
		"katoozi",
		22,
	)
	return resp, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	request := GetUserReqeust{
		ID: vars["id"],
	}
	return request, nil
}

// GetUserHandler will give you the GetUser service server
func GetUserHandler() *httptransport.Server {
	svc := service{}
	return httptransport.NewServer(
		makeGetUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
	)
}
