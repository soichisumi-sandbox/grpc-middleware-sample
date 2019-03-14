package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/soichisumi-sandbox/grpc-middleware-sample/api-pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	Users map[string]apipb.User // Todo: goroutine unsafe
}

func NewServer() (*UserServiceServer, error) {
	userDB := make(map[string]apipb.User)
	return &UserServiceServer{
		Users: userDB,
	}, nil
}

func (s *UserServiceServer) AddUser(ctx context.Context, req *apipb.AddUserRequest) (*apipb.AddUserResponse, error) {
	user := req.User
	fmt.Println("given user: ")
	spew.Println(user)
	if user.Name == "" {
		fmt.Printf("username is empty. user: %+v\n", user)
		return &apipb.AddUserResponse{}, status.Error(codes.InvalidArgument, "")
	}
	fmt.Println("s.Users: ")
	spew.Println(s.Users)
	s.Users[user.Name] = *user

	return &apipb.AddUserResponse{}, nil
}

// authorization required
func (s *UserServiceServer) GetUser(ctx context.Context, req *apipb.GetUserRequest) (*apipb.GetUserResponse, error) {
	username := req.Username
	if username == "" {
		fmt.Printf("username is empty. username: %s\n", username)
		return &apipb.GetUserResponse{}, status.Error(codes.InvalidArgument, "")
	}
	user, ok := s.Users[username]
	if !ok {
		fmt.Println("given user is not found")
		return &apipb.GetUserResponse{}, status.Error(codes.NotFound, "")
	}
	return &apipb.GetUserResponse{
		User: &user,
	}, nil
}
