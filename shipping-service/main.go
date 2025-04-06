package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "shipping-service/proto" // ganti sesuai module kamu

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedShippingServiceServer
}

func (s *server) StartShipping(ctx context.Context, req *pb.StartShippingRequest) (*pb.StartShippingResponse, error) {
	fmt.Println("Shipping started for Order ID:", req.OrderId)
	return &pb.StartShippingResponse{
		Success: true,
		Message: "Shipping started successfully",
	}, nil
}

func (s *server) CancelShipping(ctx context.Context, req *pb.CancelShippingRequest) (*pb.CancelShippingResponse, error) {
	fmt.Println("Shipping cancelled for Order ID:", req.OrderId)
	return &pb.CancelShippingResponse{
		Success: true,
		Message: "Shipping cancelled successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &server{})
	fmt.Println("Shipping Service running on port 50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
