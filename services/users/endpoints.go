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
