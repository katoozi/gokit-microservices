package users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetUserEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserReqeust)
		v, err := svc.GetUser(req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func makeUpdateUserEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		v, err := svc.UpdateUser(req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func makeDeleteUserEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		v, err := svc.DeleteUser(req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func makeCreateUserEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		v, err := svc.CreateUser(req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
