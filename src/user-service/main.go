package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	pb "example.com/user-service"
	"google.golang.org/grpc"
)

type server struct { pb.UnimplementedUserServiceServer }

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	// Simulate user retrieval
	time.Sleep(10 * time.Millisecond)
	return &pb.UserResponse{Username: "user1"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}