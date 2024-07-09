package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/grpc-example/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("client not connected: ", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetUser(ctx, &pb.GetUserRequest{Id: 1})
	if err != nil {
		fmt.Println("did not get user: ", err)
	}
	fmt.Println("User: ", r.User)
	usersResponse, err := client.ListUsers(ctx, &pb.ListUsersRequest{Ids: []int32{1, 2}})
	if err != nil {
		fmt.Println("did not get users: ", err)
	}
	fmt.Println("one and more users: ", usersResponse.Users)

	searchResponse, err := client.SearchUsers(ctx, &pb.SearchUsersRequest{City: "LA"})
	if err != nil {
		fmt.Println("did not search users: ", err)
	}
	fmt.Println("Search Users Results: ", searchResponse.Users)
}
