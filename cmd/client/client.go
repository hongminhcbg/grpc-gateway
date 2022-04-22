package main

import (
	"context"
	"fmt"

	"github.com/hongminhcbg/grpc-gateway/api"
	"google.golang.org/grpc"
)

const serverAddress = "localhost:10000"
func main() {
	conn, err := grpc.DialContext(context.Background(), serverAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cli := api.NewUserServiceClient(conn)
	resp, err := cli.CreateUser(context.Background(), &api.CreateUserRequest{Name: "Fdjskfdks", Age: 18})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
