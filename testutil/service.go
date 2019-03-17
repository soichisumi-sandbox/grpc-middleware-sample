// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

/*
Package `grpc_testing` provides helper functions for testing validators in this package.
*/

package testing

import (
	"context"
	"github.com/soichisumi-sandbox/grpc-middleware-sample/api-pb"
	"testing"
)

type TestUserService struct {
	T *testing.T
}

func (s *TestUserService) AddUser(ctx context.Context, req *apipb.AddUserRequest) (*apipb.AddUserResponse, error) {
	return &apipb.AddUserResponse{}, nil
}

func (s *TestUserService) GetUser(ctx context.Context, req *apipb.GetUserRequest) (*apipb.GetUserResponse, error) {
	return &apipb.GetUserResponse{User: &apipb.User{Name: req.Username}}, nil
}
