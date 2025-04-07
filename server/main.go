package main

import (
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Pawan2061/timeline_grpc_go/grpc"
	"github.com/Pawan2061/timeline_grpc_go/server/graph"
	"github.com/Pawan2061/timeline_grpc_go/server/graph/generated"
	"github.com/Pawan2061/timeline_grpc_go/server/store"
	grpcserver "google.golang.org/grpc"
)

func main() {
	store := store.NewStore()

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen on port 50051: %v", err)
		}

		s := grpcserver.NewServer()
		grpc.RegisterPostServiceServer(s, &grpc.PostService{Store: store})
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	resolver := graph.NewResolver(store)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("GraphQL server listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
