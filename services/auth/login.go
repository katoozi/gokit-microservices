package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// LoginRequest is the login request strucutre
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse is the login response strucutre
type LoginResponse struct {
	Token string `json:"token"`
}

func (ser service) Login(req LoginRequest) (LoginResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"password": req.Password,
	})
	// TODO: check for user exists or not
	tokenString, error := token.SignedString([]byte(ser.Secret))
	if error != nil {
		fmt.Println(error)
	}
	// TODO: save token somewhere and set the expire time for it.
	return LoginResponse{Token: tokenString}, nil
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// UpdateUserHandler will give you the UpdateUser service server
func UpdateUserHandler(secret string) *httptransport.Server {
	svc := service{}
	svc.Secret = secret
	return httptransport.NewServer(
		makeLoginEndpoint(svc),
		decodeLoginRequest,
		encodeResponse,
	)
}

func makeLoginEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		v, err := svc.Login(req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
