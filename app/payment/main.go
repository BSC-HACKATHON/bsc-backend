package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/lai0xn/bsc-payment/pkg/pb"

	"google.golang.org/grpc"
)

type paymentServiceServer struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *paymentServiceServer) GeneratePaymentLink(ctx context.Context, req *pb.GeneratePaymentLinkRequest) (*pb.GeneratePaymentLinkResponse, error) {
	link := fmt.Sprintf("https://payment.example.com/pay?user_id=%s&amount=%.2f", req.UserId, req.Amount)
	return &pb.GeneratePaymentLinkResponse{PaymentLink: link}, nil
}

func (s *paymentServiceServer) StorePayment(ctx context.Context, req *pb.StorePaymentRequest) (*pb.StorePaymentResponse, error) {
	log.Printf("Storing payment: %+v\n", req)
	return &pb.StorePaymentResponse{Success: true, Message: "Payment stored successfully"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, &paymentServiceServer{})

	log.Println("PaymentService is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
