package users

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// DeleteUserRequest is the DeleteUser service request structure
type DeleteUserRequest struct {
	ID string `json:"id"`
}

// DeleteUserResponse is the DeleteUser service response
type DeleteUserResponse struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
	Err     string `json:"error"`
}

func (ser service) DeleteUser(req DeleteUserRequest) (DeleteUserResponse, error) {
	if req.ID == "" {
		return DeleteUserResponse{}, errEmptyID
	}
	// TODO: implement the delete action(sql query for example)
	resp := DeleteUserResponse{
		ID:      req.ID,
		Deleted: true,
		Err:     "",
	}
	return resp, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request DeleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DeleteUserHandler will give you the DeleteUser service server
func DeleteUserHandler() *httptransport.Server {
	svc := service{}
	return httptransport.NewServer(
		makeDeleteUserEndpoint(svc),
		decodeDeleteUserRequest,
		encodeResponse,
	)
}
