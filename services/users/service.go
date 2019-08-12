package users

import (
	"context"
	"encoding/json"
	"net/http"
)

// Service is the users microservice interface
type Service interface {
	GetUser(GetUserReqeust) (GetUserResponse, error)
	UpdateUser(UpdateUserRequest) (UpdateUserResponse, error)
	DeleteUser(DeleteUserRequest) (DeleteUserResponse, error)
	CreateUser(CreateUserRequest) (CreateUserResponse, error)
}

type service struct{}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	// if you wanto change response code do it here.
	return json.NewEncoder(w).Encode(response)
}
