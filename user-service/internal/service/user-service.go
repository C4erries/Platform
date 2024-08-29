package service

import (
	"context"

	pb "example.com/user-service/proto"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		Id:   req.Id,
		Name: "Alice",
	}, nil
}
