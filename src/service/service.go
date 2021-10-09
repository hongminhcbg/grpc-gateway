package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hongminhcbg/grpc-gateway/api"
	"github.com/hongminhcbg/grpc-gateway/src/erp"
	"github.com/hongminhcbg/grpc-gateway/src/logger"
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
		logger.Error(fmt.Errorf("Get headers error"), "Get header error")
		return nil, erp.ErrorInternalServer
	}

	logger.Infor("Get header success", "headers", headers)
	if req.Name == "minh" {
		logger.Error(fmt.Errorf("create user error"), "create user error")
		return &api.CreateUserResponse{}, erp.ErrorBadRequest
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
