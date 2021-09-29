package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hongminhcbg/grpc-gateway/api"
	"google.golang.org/grpc/metadata"
)

type Service struct {
	api.UnimplementedUserServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		panic("get headers error")
	}

	fmt.Printf("headers = %v", headers)

	if req.Name == "minh" {
		log.Println("error")
		return &api.CreateUserResponse{}, fmt.Errorf("name existed")
	}

	log.Println("create user success")
	return &api.CreateUserResponse{
		Code:     1,
		Messsage: "success",
		Data: &api.CreateUserResponse_Data{
			UserId: time.Now().Unix(),
		},
	}, nil
}
