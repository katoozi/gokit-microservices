package users

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// GetUserResponse is the GetUser reponse structure
type GetUserResponse struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
}

// NewGetUserResponse is the GetUserResponse factory function
func NewGetUserResponse(username, firstName, lastName, password string, age int) GetUserResponse {
	return GetUserResponse{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
		Age:       age,
	}
}

// GetUserReqeust is the GetUser request structure
type GetUserReqeust struct {
	Username string `json:"username"`
}

func (ser service) GetUser(user GetUserReqeust) (GetUserResponse, error) {
	if user.Username == "" {
		return GetUserResponse{}, errEmptyUsername
	}
	// TODO: check user does exsist or not(do sql query for example).
	resp := NewGetUserResponse(
		user.Username,
		"mohammad",
		"katoozi",
		"test",
		22,
	)
	return resp, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetUserReqeust
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
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
