package main

import (
	"context"
	"log"
	"net"

	_ "github.com/99designs/gqlgen/graphql"

	pb "github.com/Pawan2061/timeline_grpc_go/grpc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPostServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	return &pb.ListPostsResponse{
		Posts: []*pb.Post{
			{
				Id:        "1",
				Content:   "Hello from gRPC!",
				Author:    "UserA",
				Timestamp: "2025-04-06T23:59:00Z",
			},
			{
				Id:        "2",
				Content:   "Hello from gRPC!",
				Author:    "UserA",
				Timestamp: "2025-04-06T23:59:00Z",
			},
			{
				Id:        "3",
				Content:   "Hello from gRPC!",
				Author:    "UserA",
				Timestamp: "2025-04-06T23:59:00Z",
			},
			{
				Id:        "4",
				Content:   "Hello from gRPC!",
				Author:    "UserA",
				Timestamp: "2025-04-06T23:59:00Z",
			},
			{
				Id:        "5",
				Content:   "Hello from gRPC!",
				Author:    "UserA",
				Timestamp: "2025-04-06T23:59:00Z",
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
