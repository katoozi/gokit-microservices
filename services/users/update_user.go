package users

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// UpdateUserRequest is the UpdateUser request structure
type UpdateUserRequest struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// UpdateUserResponse is the UpdateUser reponse strucutre
type UpdateUserResponse struct {
	UpdateUserRequest
}

func (ser service) UpdateUser(req UpdateUserRequest) (UpdateUserResponse, error) {
	if req.ID == "" {
		return UpdateUserResponse{}, errEmptyID
	}
	// TODO: do sql update query for user with req.ID
	resp := UpdateUserResponse{
		req,
	}
	resp.Username = "test"
	resp.FirstName = "test_first"
	resp.LastName = "test_last"
	resp.Age = 22
	return resp, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// UpdateUserHandler will give you the UpdateUser service server
func UpdateUserHandler() *httptransport.Server {
	svc := service{}
	return httptransport.NewServer(
		makeUpdateUserEndpoint(svc),
		decodeUpdateUserRequest,
		encodeResponse,
	)
}
