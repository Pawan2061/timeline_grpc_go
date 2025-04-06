package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Pawan2061/timeline_grpc_go/grpc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTimelineServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.TimelineRequest) (*pb.TimelineResponse, error) {
	return &pb.TimelineResponse{Message: "Hello, World! "}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTimelineServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
