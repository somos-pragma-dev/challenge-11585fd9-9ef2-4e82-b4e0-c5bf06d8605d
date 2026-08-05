package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	pb "example.com/transaction-service"
	"google.golang.org/grpc"
)

type server struct { pb.UnimplementedTransactionServiceServer }

func (s *server) ProcessTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	// Simulate transaction processing
	time.Sleep(10 * time.Millisecond)
	return &pb.TransactionResponse{Message: "Transaction processed"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransactionServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}