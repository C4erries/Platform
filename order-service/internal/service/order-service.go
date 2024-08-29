package service

import (
	"context"
	"log"

	pb "github.com/c4erries/platform/order-service/gen"
	userpb "github.com/c4erries/platform/user-service/gen"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	userClient userpb.UserServiceClient
}

func NewOrderService(userClient userpb.UserServiceClient) *OrderService {
	return &OrderService{userClient: userClient}
}

func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// Пример вызова другого микросервиса через gRPC
	userResp, err := s.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: "1"})
	if err != nil {
		return nil, err
	}

	log.Printf("Received user info: %v", userResp)

	return &pb.GetOrderResponse{
		Id:     req.Id,
		UserId: userResp.Id,
		Item:   "Laptop",
	}, nil
}
