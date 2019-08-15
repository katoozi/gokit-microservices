package users

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
)

type (
	// CreateUserRequest is the CreateUser request structure
	CreateUserRequest struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
	}

	// CreateUserResponse is the CreateUser response structure
	CreateUserResponse struct {
		ID int `json:"id"`
		CreateUserRequest
	}
)

func (srv service) CreateUser(req CreateUserRequest) (CreateUserResponse, error) {
	if req.Username == "" {
		return CreateUserResponse{}, errEmptyUsername
	}
	// TODO: create user in db
	// generate random number for id
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	resp := CreateUserResponse{
		ID:                r1.Intn(100),
		CreateUserRequest: req,
	}
	return resp, nil
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// CreateUserHandler will give you the UpdateUser service server
func CreateUserHandler() *httptransport.Server {
	svc := service{}
	return httptransport.NewServer(
		makeCreateUserEndpoint(svc),
		decodeCreateUserRequest,
		encodeResponse,
	)
}
