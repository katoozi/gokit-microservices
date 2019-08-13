package auth

import (
	"context"
	"encoding/json"
	"net/http"
)

// Service is the interface that provide authentication
type Service interface {
	Login(LoginRequest) (LoginResponse, error)
}

type service struct {
	Secret string
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	// if you wanto change response code do it here.
	return json.NewEncoder(w).Encode(response)
}
