package grpc

import (
	"context"
	"time"

	"github.com/Pawan2061/timeline_grpc_go/server/store"
)

type PostService struct {
	UnimplementedPostServiceServer
	Store *store.Store
}

func (s *PostService) ListPostsByUser(ctx context.Context, req *ListPostsRequest) (*ListPostsResponse, error) {
	storePosts := s.Store.GetUserPosts(req.UserId)
	posts := make([]*Post, len(storePosts))
	for i, post := range storePosts {
		posts[i] = &Post{
			Id:        post.ID,
			Content:   post.Content,
			Author:    post.AuthorID,
			Timestamp: post.Timestamp.Format(time.RFC3339),
		}
	}
	return &ListPostsResponse{Posts: posts}, nil
}
