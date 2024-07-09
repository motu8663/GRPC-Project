package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/grpc-example/helloworld"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

var users = []pb.User{
	{Id: 1, Fname: "Steve", City: "SA", Phone: 1234567890, Height: 5.8, Married: true},
	{Id: 2, Fname: "Nikhil", City: "LA", Phone: 852581412, Height: 5.7, Married: false},
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	for _, user := range users {
		if user.Id == req.GetId() {
			return &pb.GetUserResponse{User: &user}, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (s *server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	var response pb.ListUsersResponse
	for _, id := range req.GetIds() {
		for _, user := range users {
			if user.Id == id {
				response.Users = append(response.Users, &user)
			}
		}
	}
	return &response, nil
}

func (s *server) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	var response pb.SearchUsersResponse
	for _, user := range users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(!req.Married || user.Married == req.Married) {
			response.Users = append(response.Users, &user)
		}
	}
	return &response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to listen: %v", err)
	}
	_, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	fmt.Println("Server proceeding user data: ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: ", err)
	}
}
