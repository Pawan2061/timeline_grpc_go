package grpcservice

import (
	"context"
	"time"

	grpcserver "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	postpb "github.com/Pawan2061/timeline_grpc_go/grpc"
	"github.com/Pawan2061/timeline_grpc_go/server/store"
)

type PostService struct {
	postpb.UnimplementedPostServiceServer
	store *store.Store
}

func NewPostService(store *store.Store) *PostService {
	return &PostService{
		store: store,
	}
}

func (s *PostService) ListPostsByUser(ctx context.Context, req *postpb.ListPostsRequest) (*postpb.ListPostsResponse, error) {
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	posts := s.store.GetUserPosts(req.UserId)
	if posts == nil {
		return &postpb.ListPostsResponse{Posts: []*postpb.Post{}}, nil
	}

	response := &postpb.ListPostsResponse{
		Posts: make([]*postpb.Post, len(posts)),
	}

	for i, post := range posts {
		response.Posts[i] = &postpb.Post{
			Id:        post.ID,
			Content:   post.Content,
			Author:    post.AuthorID,
			Timestamp: post.Timestamp.Format(time.RFC3339),
		}
	}

	return response, nil
}

func RegisterService(s *grpcserver.Server, store *store.Store) {
	postpb.RegisterPostServiceServer(s, NewPostService(store))
}
