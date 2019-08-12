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
}

type service struct{}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
