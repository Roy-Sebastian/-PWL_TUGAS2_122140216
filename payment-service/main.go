package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "payment-service/proto" 

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *server) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	fmt.Println("Payment processed for Order ID:", req.OrderId)
	return &pb.ProcessPaymentResponse{
		Success: true,
		Message: "Payment successful",
	}, nil
}

func (s *server) RefundPayment(ctx context.Context, req *pb.RefundPaymentRequest) (*pb.RefundPaymentResponse, error) {
	fmt.Println("Payment refunded for Order ID:", req.OrderId)
	return &pb.RefundPaymentResponse{
		Success: true,
		Message: "Payment refunded successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})
	fmt.Println("Payment Service running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
